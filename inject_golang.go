package main

import (
	`os`
	`path/filepath`

	`github.com/storezhang/simaqian`
)

func golang(conf *config, output string, logger simaqian.Logger) error {
	return filepath.Walk(output, func(path string, info os.FileInfo, walkErr error) (err error) {
		if nil != walkErr {
			err = walkErr
		}
		if nil != err {
			return
		}
		if info.IsDir() {
			return
		}

		if matched, matchErr := filepath.Match(`*.pb.go`, filepath.Base(path)); matchErr != nil {
			err = matchErr
		} else if matched {
			err = gtag(conf, path, logger)
		}

		return
	})
}
