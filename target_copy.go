package main

import (
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

func (t *target) copy(source string, logger simaqian.Logger, filenames ...string) (err error) {
	for _, filename := range filenames {
		var needs []string
		if needs, err = gfx.All(source, gfx.Pattern(filename)); nil != err {
			continue
		}

		// 对列出的所有文件逐一复制
		for _, need := range needs {
			destFilename := gfx.Name(need, gfx.File(), gfx.Ext(filepath.Ext(need)))
			to := filepath.Join(t.output(), destFilename)
			// 目的文件已经存在，不应该执行文件复制
			if _, exists := gfx.Exists(to); exists {
				continue
			}

			if err = gfx.Copy(need, to); nil != err {
				logger.Error("复制文件出错", field.New("from", need), field.New("to", to), field.Error(err))
			}
		}
	}

	return
}
