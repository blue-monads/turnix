package hookengine

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/bornjre/turnix/backend/xtypes/services/xhook"
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

var regex = regexp.MustCompile(`const handle = \((.*?)\) => {`)

var (
	ErrMustRun = errors.New("COULD NOT RUN")
)

type parsedHook struct {
	id          int64
	name        string
	eventType   string
	hookType    string
	runasUserID int64
	envs        map[string]string
	target      string
	mustRun     bool
	mustNoError bool
}

type hookRunner struct {
	parent *HookEngine
	pid    int64

	jsCodeCache *goja.Program

	parsedHooks []parsedHook
}

func newHookRunner(h *HookEngine, pid int64, hooks []models.ProjectHook) (*hookRunner, error) {

	h.logger.Debug().Int64("pid", pid).Any("hooks", hooks).Msg("newHookRunner")

	var codeBuf strings.Builder

	for _, hook := range hooks {
		if hook.HookType != "script" {
			h.logger.Info().Int64("hook_id", hook.ID).Msg("newHookRunner/skipping_not_script")
			continue
		}

		updatedCode := regex.ReplaceAllString(hook.HookCode, fmt.Sprintf("const _handle_%d = ($1) => {", hook.ID))
		codeBuf.WriteString(updatedCode)
	}

	code := codeBuf.String()

	h.logger.Debug().Str("code", code).Msg("newHookRunner/finalCode")

	program, err := goja.Compile(fmt.Sprintf("project_hook_%d", pid), code, true)
	if err != nil {
		h.logger.Warn().Str("code", code).Msg("newHookRunner/goja_compile_error")
		return nil, err
	}

	parsedHooks := make([]parsedHook, 0, len(hooks))

	for _, hook := range hooks {

		envs := make(map[string]string)
		err = json.Unmarshal([]byte(hook.Envs), &envs)
		if err != nil {
			h.logger.Warn().Str("env", hook.Envs).Msg("newHookRunner/decoding_env_error")
		}

		parsedHooks = append(parsedHooks, parsedHook{
			id:          hook.ID,
			name:        hook.Name,
			eventType:   hook.Event,
			hookType:    hook.HookType,
			runasUserID: hook.RunasUserID,
			envs:        envs,
			target:      hook.Target,
			mustRun:     false,
			mustNoError: false,
		})

	}

	return &hookRunner{
		parent:      h,
		pid:         pid,
		jsCodeCache: program,
		parsedHooks: parsedHooks,
	}, nil

}

func (r *hookRunner) execute(evt xhook.Event) (*xhook.Result, error) {

	gojah := r.parent.gojaPool.Get(evt.ProjectId, evt.Id, false)
	if gojah == nil {
		return nil, fmt.Errorf("could not accure JS runtime")
	}

	if gojah.lastPid != evt.ProjectId {
		gojah.lastPid = 0
		_, err := gojah.js.RunProgram(r.jsCodeCache)
		if err != nil {
			return nil, err
		}

		gojah.bind()
		gojah.js.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

		gojah.lastPid = evt.ProjectId

	}

	execCtx := Executor{
		Event:         evt,
		JsRuntime:     gojah.js,
		PreventAction: false,
		ResultData:    map[string]any{},
	}

	result := &xhook.Result{
		NoOfHooksRan:  0,
		PreventAction: false,
		Errors:        map[string]string{},
		Data:          nil,
	}

	defer func() {
		result.Data = execCtx.ResultData
	}()

	for _, ph := range r.parsedHooks {

		pp.Println("@ph", ph)

		r.parent.logger.Info().Any("parsed_hook", ph).Msg("execute")

		if ph.eventType != evt.Type {
			continue
		}

		var err error
		var ran = true

		switch ph.hookType {
		case "script":
			pp.Println("@@@@@")
			err = execCtx.executeJS(ph)
		case "webhook":
			err = execCtx.executeWebhook(ph)
		default:
			ran = false
			log.Println("unknown_hookType", ph.hookType)
		}

		if ph.mustRun && !ran {
			return nil, ErrMustRun
		}

		if ran {
			result.NoOfHooksRan = result.NoOfHooksRan + 1
		}

		if err != nil {
			if ph.mustNoError {
				return nil, err
			}

			result.Errors[ph.name] = err.Error()
		}

		if execCtx.PreventAction {
			return result, nil
		}

	}

	return result, nil
}
