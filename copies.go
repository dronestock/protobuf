package main

import (
	`path/filepath`

	`github.com/goexl/gfx`
	`github.com/storezhang/gox/field`
)

func (p *plugin) copies() (undo bool, err error) {
	for _, output := range p.Outputs {
		for _, filename := range p.Copies {
			from := filepath.Join(p.Source, filename)
			// 检查源文件是否存在，不存在复制下一个文件
			if !gfx.Exists(from) {
				continue
			}

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
