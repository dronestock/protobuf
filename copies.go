package main

import (
	`path/filepath`

	`github.com/storezhang/gfx`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) copies(logger simaqian.Logger) (undo bool, err error) {
	for lang, inputs := range p.config.inputsCache {
		for _, input := range inputs {
			for _, _copy := range p.config.Copies {
				from := filepath.Join(input, _copy)
				to := filepath.Join(p.config.outputCache[lang], _copy)
				if err = gfx.Copy(from, to); nil != err {
					logger.Error(`复制文件出错`, field.String(`from`, from), field.String(`to`, to), field.Error(err))
				}
				if nil != err {
					return
				}
			}
		}
	}

	return
}
