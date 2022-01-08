package main

import (
	`github.com/storezhang/simaqian`
)

func inject(conf *config, logger simaqian.Logger) (err error) {
	for _type, output := range conf.outputCache {
		switch _type {
		case langGo:
			fallthrough
		case langGogo:
			err = golang(conf, output, logger)
		}
	}

	return
}
