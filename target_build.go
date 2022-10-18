package main

import (
	"fmt"

	"github.com/goexl/gfx"
)

func (t *target) build(plugin *plugin) (err error) {
	args := []interface{}{
		// 加入当前目录
		// 防止出现错误：File does not reside within any path specified using --proto_path
		`--proto_path`, plugin.Source,
	}

	// 添加导入目录
	for _, include := range plugin.Includes {
		args = append(args, `--proto_path`, include)
	}

	// 添加标签
	for _, tag := range plugin.tags() {
		args = append(args, fmt.Sprintf(`--%s`, tag))
	}

	// 添加插件他输出目录
	args = append(args, t.out(plugin.Defaults))
	// 添加选项
	args = append(args, t.opt())

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
