package hookengine

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/bornjre/trunis/backend/xtypes/models"
	"github.com/dop251/goja"
)

var regex = regexp.MustCompile(`const handle = \((.*?)\) => {`)

type parsedHook struct {
	id          int64
	hookType    string
	runasUserID int64
	envs        map[string]string
}

type hookRunner struct {
	parent      *HookEngine
	pid         int64
	jsCodeCache *goja.Program

	parsedHooks []parsedHook

	compileError bool
	message      string
}

func newHookRunner(h *HookEngine, pid int64, hooks []models.ProjectHook) *hookRunner {

	var codeBuf strings.Builder

	for _, hook := range hooks {
		if hook.HookType != "script" {
			continue
		}

		updatedCode := regex.ReplaceAllString(hook.HookCode, fmt.Sprintf("const _handle_%d = ($1) => {", hook.ID))
		codeBuf.WriteString(updatedCode)
	}

	program, err := goja.Compile(fmt.Sprintf("project_hook_%d", pid), codeBuf.String(), true)
	if err != nil {
		return &hookRunner{
			compileError: true,
			message:      err.Error(),
		}

	}

	parsedHooks := make([]parsedHook, 0, len(hooks))

	for _, hook := range hooks {

		envs := make(map[string]string)
		err = json.Unmarshal([]byte(hook.Envs), &envs)
		if err != nil {
			log.Println("decoding_env_error", err)
		}

		parsedHooks = append(parsedHooks, parsedHook{
			id:          hook.ID,
			hookType:    hook.HookType,
			runasUserID: hook.RunasUserID,
			envs:        envs,
		})

	}

	return &hookRunner{
		parent:      h,
		pid:         pid,
		jsCodeCache: program,
		parsedHooks: parsedHooks,
	}

}

func (r *hookRunner) execute(evt xtypes.HookEvent) (result xtypes.HookResult) {

	gojah := r.parent.gojaPool.Get(evt.ProjectId, false)
	if gojah == nil {
		result.Error = fmt.Errorf("could not accure JS runtime")
		return
	}

	if gojah.lastPid != evt.ProjectId {

		_, err := gojah.js.RunProgram(r.jsCodeCache)
		if err != nil {
			result.Error = err
			return
		}

	}

	for _, ph := range r.parsedHooks {

		switch ph.hookType {
		case "script":
			result.NoOfHooksRan = result.NoOfHooksRan + 1
			return r.executeJS(evt, ph, gojah)
		case "webhook":
			result.NoOfHooksRan = result.NoOfHooksRan + 1
			return r.executeWebhook(evt, ph)
		default:
			log.Println("unknown_hookType", ph.hookType)
		}

	}

	return
}

func (r *hookRunner) executeWebhook(evt xtypes.HookEvent, hook parsedHook) (result xtypes.HookResult) {

	return

}

func (r *hookRunner) executeJS(evt xtypes.HookEvent, hook parsedHook, jhandle *gojaHandle) (result xtypes.HookResult) {

	funcName := fmt.Sprintf("_handle_%d", hook.id)

	var entry func(ctx *goja.Object)
	eval := jhandle.js.Get(funcName)
	if eval == nil {
		result.Error = fmt.Errorf("%s function not found in script", funcName)
		return
	}

	err := jhandle.js.ExportTo(eval, &entry)
	if err != nil {
		result.Error = err

	}

	obj := buildEventObject(evt, jhandle.js)

	entry(obj)

	return

}

func buildEventObject(evt xtypes.HookEvent, r *goja.Runtime) *goja.Object {
	obj := r.NewObject()

	obj.Set("dataAsObject", func() any {

		return evt.Data
	})

	obj.Set("name", evt.Name)

	return obj

}
