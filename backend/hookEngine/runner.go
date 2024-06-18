package hookengine

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bornjre/trunis/backend/xtypes/models"
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

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

func (h *HookEngine) newHookRunner(pid int64, hooks []models.ProjectHook) *hookRunner {

	var codeBuf strings.Builder

	for _, hook := range hooks {
		if hook.HookType != "script" {
			continue
		}

		codeBuf.WriteString(hook.HookCode)
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
		parent:       h,
		pid:          pid,
		jsCodeCache:  program,
		compileError: false,
		message:      "",
		parsedHooks:  parsedHooks,
	}

}

func (r *hookRunner) execute() error {

	for _, ph := range r.parsedHooks {
		pp.Println(ph)

	}

	return nil
}
