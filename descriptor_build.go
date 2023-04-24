package main

import (
	"path/filepath"

	"github.com/goexl/gox/args"
)

func (d *descriptor) build(plugin *plugin) (err error) {
	ba := args.New().Build()
	// 添加导入目录
	for _, include := range plugin.Includes {
		if abs, ae := filepath.Abs(include); nil != ae {
			err = ae
		} else {
			ba.Option("proto_path", abs)
		}

		if nil != err {
			return
		}
	}

	// 添加选项
	for _, opt := range d.Opts {
		ba.Add(opt)
	}
	// 包含导入文件
	ba.Flag("include_imports")
	// 添加输出文件
	ba.Option("descriptor_set_out", d.Output)
	// 编译
	err = plugin.protoc(plugin.Source, append([]string{d.Source}, d.Sources...), ba)

	return
}
