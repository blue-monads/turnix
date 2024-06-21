package hookengine

import (
	"errors"

	"github.com/bornjre/turnix/backend/xtypes"
)

var (
	ErrRuntimeErr       = errors.New("RUNTIME ERROR")
	ErrScriptCompileErr = errors.New("SCRIPT COMPILE ERROR")
)

func (h *HookEngine) emit(evt xtypes.HookEvent) (*xtypes.HookResult, error) {

	runner := h.getRunner(evt.ProjectId)

	if runner == nil {
		return nil, ErrRuntimeErr
	}

	if runner.compileError {
		return nil, ErrScriptCompileErr
	}

	return runner.execute(evt)
}

func (h *HookEngine) getRunner(pid int64) *hookRunner {

	h.hrLock.RLock()
	runner := h.hookRunners[pid]
	if runner != nil {
		h.hrLock.RUnlock()
		return runner
	}

	h.hrLock.RUnlock()

	hooks, err := h.db.ListProjectHooks(pid)
	if err != nil {

		return nil
	}

	runner = newHookRunner(h, pid, hooks)

	// race condition is fine

	h.hrLock.Lock()
	defer h.hrLock.Unlock()

	racedrunner := h.hookRunners[pid]
	if runner != nil {
		return racedrunner
	}

	h.hookRunners[pid] = runner

	return runner

}
