package main

func (p *plugin) inject() (undo bool, err error) {
	for _, _target := range p.Targets {
		if err = _target.inject(p); nil != err {
			return
		}
	}

	return
}
