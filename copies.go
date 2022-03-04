package main

import (
	`path/filepath`

	`github.com/storezhang/gfx`
	`github.com/storezhang/gox/field`
)

func (p *plugin) copies() (undo bool, err error) {
	for lang, inputs := range p.inputsCache {
		for _, input := range inputs {
			for _, _copy := range p.Copies {
				from := filepath.Join(input, _copy)
				to := filepath.Join(p.outputCache[lang], _copy)
				if err = gfx.Copy(from, to); nil != err {
					p.Error(`复制文件出错`, field.String(`from`, from), field.String(`to`, to), field.Error(err))
				}
				if nil != err {
					return
				}
			}
		}
	}

	return
}
