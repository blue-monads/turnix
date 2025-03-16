package engine

func (e *Engine) OnInit(name string, id int64) error {
	e.pLock.RLock()
	def := e.projects[name]
	e.pLock.RUnlock()

	err := def.def.OnInit(id)
	if err != nil {
		return err
	}

	return nil

}
