package main

import (
	`github.com/storezhang/simaqian`
)

func builds(conf *config, logger simaqian.Logger) (err error) {
	for lang, inputs := range conf.inputsCache {
		for _, input := range inputs {
			if err = build(conf, lang, input, conf.outputCache[lang], logger); nil != err {
				return
			}
		}
	}

	return
}
