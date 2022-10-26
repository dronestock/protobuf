package main

import (
	"fmt"
	"path/filepath"

	"github.com/goexl/gfx"
)

func (t *target) build(plugin *plugin) (err error) {
	args := make([]any, 0, 16)

	// 加入当前目录
	// 防止出现错误：File does not reside within any path specified using --proto_path
	if abs, ae := filepath.Abs(plugin.Source); nil != ae {
		err = ae
	} else {
		args = append(args, `--proto_path`, abs)
	}
	if nil != err {
		return
	}

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

	// 添加标签
	for _, tag := range plugin.tags() {
		args = append(args, fmt.Sprintf(`--%s`, tag))
	}

	// 添加插件他输出目录
	for _, out := range t.out(plugin.Defaults) {
		args = append(args, out)
	}
	// 添加选项
	for _, opt := range t.opt(plugin.Defaults) {
		args = append(args, opt)
	}

	// 编译
	if filenames, ge := gfx.All(plugin.Source, gfx.Matchable(plugin.buildable)); nil != ge {
		err = ge
	} else {
		for _, filename := range filenames {
			if err = plugin.protoc(plugin.Source, filename, args...); nil != err {
				break
			}
		}
	}

	return
}
