package main

import (
	"fmt"

	"github.com/goexl/gfx"
)

func (t *target) build(plugin *plugin) (err error) {
	args := []interface{}{
		// 加入当前目录
		// 防止出现错误：File does not reside within any path specified using --proto_path
		fmt.Sprintf(`--proto_path=%s`, plugin.Source),
	}

	// 添加导入目录
	for _, include := range plugin.Includes {
		args = append(args, fmt.Sprintf(`--proto_path=%s`, include))
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
	walkOptions := gfx.NewWalkOptions(
		gfx.Pattern(protoFilePattern),
		gfx.Matchable(plugin.buildable),
	)
	if filenames, ge := gfx.All(plugin.Source, walkOptions...); nil != ge {
		err = ge
	} else {
		for _, filename := range filenames {
			if ge = plugin.protoc(plugin.Source, filename, args...); nil != err {
				break
			}
		}
	}

	return
}
