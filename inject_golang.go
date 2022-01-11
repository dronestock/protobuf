package main

import (
	`github.com/storezhang/gfx`
	`github.com/storezhang/simaqian`
)

func golang(conf *config, output string, logger simaqian.Logger) (err error) {
	var files []string
	if files, err = gfx.All(output, gfx.Pattern(conf.protoGoFilePattern)); nil != err {
		return
	}

	for _, file := range files {
		if err = gtag(conf, file, logger); nil != err {
			break
		}
	}

	return
}
