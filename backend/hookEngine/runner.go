package hookengine

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bornjre/trunis/backend/xtypes/models"
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
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

type EventContext struct {
	UserId    int64
	ProjectId int64
	EventId   int64
}

func (r *hookRunner) execute(evt EventContext) error {

	for _, ph := range r.parsedHooks {

		funcName := fmt.Sprintf("_handle_%d", ph.id)

		pp.Println("EXECUTING", funcName)

	}

	return nil
}
