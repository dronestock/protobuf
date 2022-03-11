package main

func (p *plugin) inject(typ string) (err error) {
	switch typ {
	case typeGo:
		fallthrough
	case typeGogo:
		err = p.golang(p.output(typ))
	}

	return
}
