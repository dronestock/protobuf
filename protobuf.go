package main

import (
	`github.com/storezhang/gfx`
	`github.com/storezhang/simaqian`
)

func (p *plugin) protobuf(input string, args []string, logger simaqian.Logger) (err error) {
	var files []string
	if files, err = gfx.All(input, gfx.Pattern(protoFilePattern), gfx.Matchable(p.config.buildable)); nil != err {
		return
	}

	for _, file := range files {
		if err = p.protoc(file, logger, args...); nil != err {
			break
		}
	}

	return
}
