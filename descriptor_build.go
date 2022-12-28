package main

import (
	"path/filepath"
)

func (d *descriptor) build(plugin *plugin) (err error) {
	args := make([]any, 0, 4)
	// 添加导入目录
	for _, include := range plugin.Includes {
		if abs, ae := filepath.Abs(include); nil != ae {
			err = ae
		} else {
			args = append(args, `--proto_path`, abs)
		}

		if nil != err {
			return
		}
	}

	// 添加选项
	for _, opt := range d.Opts {
		args = append(args, opt)
	}
	// 包含导入文件
	args = append(args, "--include_imports")

	// 添加输出文件
	args = append(args, "--descriptor_set_out", d.Output)

	// 编译
	err = plugin.protoc(plugin.Source, append([]string{d.Source}, d.Sources...), args)

	return
}
