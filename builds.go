package main

func (p *plugin) builds() (undo bool, err error) {
	for lang := range p.Outputs {
		if err = p.build(lang); nil != err {
			return
		}
	}

	return
}
