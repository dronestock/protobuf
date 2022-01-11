package main

import (
	`github.com/storezhang/gfx`
	`github.com/storezhang/simaqian`
)

func protobuf(conf *config, input string, args []string, logger simaqian.Logger) (err error) {
	var files []string
	if files, err = gfx.All(input, gfx.Pattern(conf.protoFilePattern), gfx.Matchable(conf.buildable)); nil != err {
		return
	}

	for _, file := range files {
		if err = protoc(conf, file, logger, args...); nil != err {
			break
		}
	}

	return
}
