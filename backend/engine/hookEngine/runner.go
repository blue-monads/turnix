package hookengine

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/bornjre/turnix/backend/xtypes/xengine"
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

var regex = regexp.MustCompile(`const handle = \((.*?)\) => {`)

var (
	ErrMustRun = errors.New("COULD NOT RUN")
)

type hookRunner struct {
	parent *HookEngine
	pid    int64

	jsCodeCache *goja.Program

	parsedHooks []ParsedHook
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

	parsedHooks := make([]ParsedHook, 0, len(hooks))

	for _, hook := range hooks {

		envs := make(map[string]string)
		err = json.Unmarshal([]byte(hook.Envs), &envs)
		if err != nil {
			h.logger.Warn().Str("env", hook.Envs).Msg("newHookRunner/decoding_env_error")
		}

		parsedHooks = append(parsedHooks, ParsedHook{
			Id:          hook.ID,
			Name:        hook.Name,
			EventType:   hook.Event,
			HookType:    hook.HookType,
			RunasUserID: hook.RunasUserID,
			Envs:        envs,
			Target:      hook.Target,
			MustRun:     false,
			MustNoError: false,
		})

	}

	return &hookRunner{
		parent:      h,
		pid:         pid,
		jsCodeCache: program,
		parsedHooks: parsedHooks,
	}, nil

}

func (r *hookRunner) execute(evt xengine.EventContext) (*xengine.EventResult, error) {

	gojah := r.parent.gojaPool.Get(evt.Event.Project, false)
	if gojah == nil {
		return nil, fmt.Errorf("could not accure JS runtime")
	}

	if gojah.LastPid != evt.Event.Project {
		gojah.LastPid = 0
		_, err := gojah.JS.RunProgram(r.jsCodeCache)
		if err != nil {
			return nil, err
		}

		gojah.Bind()
		gojah.JS.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

		gojah.LastPid = evt.Event.Project

	}

	execCtx := Executor{
		engine:        r.parent,
		Event:         evt,
		JsRuntime:     gojah.JS,
		PreventAction: false,
		ResultData:    map[string]any{},
	}

	result := &xengine.EventResult{
		PreventAction: false,
		Errors:        map[string]string{},
		Data:          nil,
	}

	noHook := 0

	defer func() {
		r.parent.gojaPool.Set(gojah)

		result.Data = execCtx.ResultData
	}()

	for _, ph := range r.parsedHooks {

		pp.Println("@ph", ph)

		r.parent.logger.Info().Any("parsed_hook", ph).Msg("execute")

		if ph.EventType != evt.Event.Type {
			continue
		}

		var err error
		var ran = true

		switch ph.HookType {
		case "script":
			pp.Println("@@@@@")
			err = execCtx.executeJS(ph)
		case "webhook":
			err = execCtx.executeWebhook(ph)
		default:
			ran = false
			log.Println("unknown_hookType", ph)
		}

		if ph.MustRun && !ran {
			return nil, ErrMustRun
		}

		if ran {
			noHook = noHook + 1
		}

		if err != nil {
			if ph.MustNoError {
				return nil, err
			}

			result.Errors[ph.Name] = err.Error()
		}

		if execCtx.PreventAction {
			return result, nil
		}

	}

	pp.Println("@ran", noHook)

	return result, nil
}
