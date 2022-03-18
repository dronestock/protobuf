package main

import (
	`path/filepath`

	`github.com/goexl/gfx`
	`github.com/goexl/gox/field`
)

func (p *plugin) copies() (undo bool, err error) {
	for _, output := range p.Outputs {
		for _, _copy := range p._copies() {
			var needs []string
			if needs, err = gfx.All(p.Source, gfx.Pattern(_copy)); nil != err {
				continue
			}

			// 对列出的所有文件逐一复制
			for _, need := range needs {
				filename := gfx.Filename(need, gfx.File())
				to := filepath.Join(output, filename)
				// 目的文件已经存在，不应该执行文件复制
				if _, exists := gfx.Exists(to); exists {
					continue
				}

				if err = gfx.Copy(need, to); nil != err {
					p.Error(`复制文件出错`, field.String(`from`, need), field.String(`to`, to), field.Error(err))
				}
			}
		}
	}

	return
}
