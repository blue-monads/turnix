package hookengine

import (
	"github.com/bornjre/turnix/backend/xtypes"
)

func (h *HookEngine) emit(evt xtypes.HookEvent) (*xtypes.HookResult, error) {

	runner := h.getRunner(evt.ProjectId)

	if runner == nil {
		return nil, nil
	}

	if runner.compileError {
		return nil, nil
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
