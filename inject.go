package main

import (
	`github.com/storezhang/simaqian`
)

func (p *plugin) inject(logger simaqian.Logger) (undo bool, err error) {
	for _type, output := range p.config.outputCache {
		switch _type {
		case langGo:
			fallthrough
		case langGogo:
			err = p.golang(output, logger)
		}
	}

	return
}
