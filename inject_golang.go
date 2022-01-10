package main

import (
	`github.com/storezhang/gox/file`
	`github.com/storezhang/simaqian`
)

func golang(conf *config, output string, logger simaqian.Logger) (err error) {
	var paths []string
	if paths, err = file.Files(output, file.Pattern(conf.protoGoFilePattern)); nil != err {
		return
	}

	for _, path := range paths {
		if err = gtag(conf, path, logger); nil != err {
			break
		}
	}

	return
}
