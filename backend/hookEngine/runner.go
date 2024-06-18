package hookengine

import (
	"fmt"
	"strings"

	"github.com/bornjre/trunis/backend/xtypes/models"
	"github.com/dop251/goja"
)

type hookRunner struct {
	parent      *HookEngine
	pid         int64
	jsCodeCache *goja.Program

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

	return &hookRunner{
		parent:       h,
		pid:          pid,
		jsCodeCache:  program,
		compileError: false,
		message:      "",
	}

}

func (run *hookRunner) execute() error {

	run.parent.gojaPool.Get(1, true)

	return nil
}
