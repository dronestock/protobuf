package main

func (p *plugin) inject(lang string) (err error) {
	switch lang {
	case langGo:
		fallthrough
	case langGogo:
		err = p.golang(p.output(lang))
	}

	return
}
