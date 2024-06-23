package hookengine

import (
	"errors"

	"github.com/bornjre/turnix/backend/xtypes/services/xhook"
)

var (
	ErrRuntimeErr       = errors.New("RUNTIME ERROR")
	ErrScriptCompileErr = errors.New("SCRIPT COMPILE ERROR")
)

func (h *HookEngine) emit(evt xhook.Event) (*xhook.Result, error) {

	runner, err := h.getRunner(evt.ProjectId)
	if err != nil {
		return nil, err
	}

	return runner.execute(evt)
}

func (h *HookEngine) getRunner(pid int64) (*hookRunner, error) {

	h.hrLock.RLock()
	runner := h.hookRunners[pid]
	if runner != nil {
		h.hrLock.RUnlock()
		return runner, nil
	}

	h.hrLock.RUnlock()

	hooks, err := h.db.ListProjectHooks(pid)
	if err != nil {
		return nil, err
	}

	runner, err = newHookRunner(h, pid, hooks)
	if err != nil {
		return nil, err
	}

	// race condition is fine

	h.hrLock.Lock()
	defer h.hrLock.Unlock()

	racedrunner := h.hookRunners[pid]
	if runner != nil {
		return racedrunner, nil
	}

	h.hookRunners[pid] = runner

	return runner, nil

}
