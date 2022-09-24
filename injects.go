package main

func (p *plugin) injects() (undo bool, err error) {
	for _, _target := range p.Targets {
		if err = _target.inject(p); nil != err {
			return
		}
	}

	return
}
