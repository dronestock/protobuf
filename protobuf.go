package main

import (
	`github.com/storezhang/gfx`
)

func (p *plugin) protobuf(input string, args []interface{}) (err error) {
	var files []string
	if files, err = gfx.All(input, gfx.Pattern(protoFilePattern), gfx.Matchable(p.buildable)); nil != err {
		return
	}

	for _, file := range files {
		if err = p.protoc(file, args...); nil != err {
			break
		}
	}

	return
}
