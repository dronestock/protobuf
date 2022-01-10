package main

import (
	`github.com/storezhang/gox/file`
	`github.com/storezhang/simaqian`
)

func protobuf(conf *config, input string, args []string, logger simaqian.Logger) (err error) {
	var paths []string
	if paths, err = file.Files(input, file.Pattern(conf.protoFilePattern), file.Matchable(conf.buildable)); nil != err {
		return
	}

	for _, path := range paths {
		if err = protoc(conf, path, logger, args...); nil != err {
			break
		}
	}

	return
}
