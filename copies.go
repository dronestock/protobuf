package main

import (
	`path/filepath`

	`github.com/storezhang/gfx`
	`github.com/storezhang/gox/field`
)

func (p *plugin) copies() (undo bool, err error) {
	for _, output := range p.Outputs {
		for _, filename := range p.Copies {
			from := filepath.Join(p.Source, filename)
			to := filepath.Join(output, filename)
			if err = gfx.Copy(from, to); nil != err {
				p.Error(`复制文件出错`, field.String(`from`, from), field.String(`to`, to), field.Error(err))
			}
			if nil != err {
				return
			}
		}
	}

	return
}
