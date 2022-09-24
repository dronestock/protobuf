package main

func (p *plugin) builds() (undo bool, err error) {
	for _, _target := range p.Targets {
		if err = _target.build(p); nil != err {
			return
		}
	}

	return
}
