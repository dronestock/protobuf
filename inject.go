package main

func (p *plugin) inject(lang string, input string) (err error) {
	switch lang {
	case langGo:
		fallthrough
	case langGogo:
		err = p.golang(input, p.output(lang))
	}

	return
}
