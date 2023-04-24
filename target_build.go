package main

import (
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
)

func (t *target) build(plugin *plugin) (err error) {
	ba := args.New().Build()

	// 加入当前目录
	// 防止出现错误：File does not reside within any path specified using --proto_path
	if abs, ae := filepath.Abs(plugin.Source); nil != ae {
		err = ae
	} else {
		ba.Option("proto_path", abs)
	}
	if nil != err {
		return
	}

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

	// 添加标签
	for _, tag := range plugin.tags() {
		ba.Flag(tag)
	}

	// 添加插件他输出目录
	for _, out := range t.out(plugin) {
		ba.Add(out)
	}
	// 添加选项
	for _, opt := range t.opt(plugin) {
		ba.Add(opt)
	}

	// 编译
	if filenames, ge := gfx.All(plugin.Source, gfx.Matchable(plugin.buildable)); nil != ge {
		err = ge
	} else {
		for _, filename := range filenames {
			if err = plugin.protoc(plugin.Source, []string{filename}, ba); nil != err {
				break
			}
		}
	}

	return
}
