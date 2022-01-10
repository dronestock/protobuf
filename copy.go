package main

import (
	`path/filepath`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/gox/file`
	`github.com/storezhang/simaqian`
)

func copies(conf *config, logger simaqian.Logger) (err error) {
	for lang, inputs := range conf.inputsCache {
		for _, input := range inputs {
			for _, _copy := range conf.Copies {
				from := filepath.Join(input, _copy)
				to := filepath.Join(conf.outputCache[lang], _copy)
				if err = file.Copy(from, to); nil != err {
					logger.Error(`复制文件出错`, field.String(`from`, from), field.String(`to`, to))
				}
				if nil != err {
					return
				}
			}
		}
	}

	return
}
