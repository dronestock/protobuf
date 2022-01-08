package main

import (
	`os`
	`path/filepath`

	`github.com/storezhang/simaqian`
)

func protobuf(conf *config, input string, args []string, logger simaqian.Logger) error {
	return filepath.Walk(input, func(path string, info os.FileInfo, walkErr error) (err error) {
		if nil != walkErr {
			err = walkErr
		}
		if nil != err {
			return
		}
		if info.IsDir() {
			return
		}

		if matched, matchErr := filepath.Match(conf.protoFilePattern, filepath.Base(path)); matchErr != nil {
			err = matchErr
		} else if matched {
			err = protoc(conf, path, logger, args...)
		}

		return
	})
}
