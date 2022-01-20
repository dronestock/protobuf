package main

import (
	`github.com/storezhang/simaqian`
)

func (p *plugin) builds(logger simaqian.Logger) (undo bool, err error) {
	for lang, inputs := range p.config.inputsCache {
		for _, input := range inputs {
			if err = p.build(lang, input, p.config.outputCache[lang], logger); nil != err {
				return
			}
		}
	}

	return
}
