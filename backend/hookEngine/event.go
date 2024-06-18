package hookengine

import (
	"errors"

	"github.com/bornjre/trunis/backend/xtypes"
)

func (h *HookEngine) emit(evt xtypes.HookEvent) xtypes.HookResult {

	runner := h.getRunner(evt.ProjectId)

	if runner != nil {
		return xtypes.HookResult{
			Error: errors.New("Runner not found"),
		}
	}

	if runner.compileError {
		return xtypes.HookResult{
			Error: errors.New(runner.message),
		}
	}

	runner.execute()

	return xtypes.HookResult{
		NoOfHooksRan: 0,
		Mutated:      false,
		Error:        nil,
	}
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

	runner = h.newHookRunner(pid, hooks)

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