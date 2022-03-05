package main

import (
	`github.com/storezhang/gfx`
)

func (p *plugin) protobuf(input string, args []interface{}) (err error) {
	var filenames []string
	if filenames, err = gfx.All(input, gfx.Pattern(protoFilePattern), gfx.Matchable(p.buildable)); nil != err {
		return
	}

	for _, filename := range filenames {
		if err = p.protoc(filename, args...); nil != err {
			break
		}
	}

	return
}
