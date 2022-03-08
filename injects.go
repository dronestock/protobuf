package main

func (p *plugin) injects() (undo bool, err error) {
	for lang := range p.Outputs {
		if err = p.inject(lang); nil != err {
			return
		}
	}

	return
}
