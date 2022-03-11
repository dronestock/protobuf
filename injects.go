package main

func (p *plugin) injects() (undo bool, err error) {
	for typ := range p.Outputs {
		if err = p.inject(typ); nil != err {
			return
		}
	}

	return
}
