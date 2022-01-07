package main

import (
	`github.com/storezhang/simaqian`
)

func build(conf *config, logger simaqian.Logger) (err error) {
	for _type, inputs := range conf.inputsCache {
		for _, input := range inputs {
			if err = protoc(conf, _type, input, conf.outputCache[_type], logger); nil != err {
				return
			}
		}
	}

	return
}
