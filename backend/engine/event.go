package engine

import "fmt"

func (e *Engine) OnInit(name string, id int64) error {
	e.pLock.RLock()
	def := e.projects[name]
	e.pLock.RUnlock()

	if def == nil {
		return fmt.Errorf("project type %s not found", name)
	}

	if def.def.OnInit != nil {
		err := def.def.OnInit(id)
		if err != nil {
			return err
		}
	}

	return nil

}
