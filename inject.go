package main

func (p *plugin) inject() (undo bool, err error) {
	for _type, output := range p.outputCache {
		switch _type {
		case langGo:
			fallthrough
		case langGogo:
			err = p.golang(output)
		}
	}

	return
}
