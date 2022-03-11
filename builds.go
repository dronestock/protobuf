package main

func (p *plugin) builds() (undo bool, err error) {
	for typ := range p.Outputs {
		if err = p.build(typ); nil != err {
			return
		}
	}

	return
}
