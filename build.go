package main

import (
	`fmt`
)

func (p *plugin) build(lang string) (err error) {
	args := []interface{}{
		// 加入当前目录
		// 防止出现错误：File does not reside within any path specified using --proto_path
		fmt.Sprintf(`--proto_path=%s`, p.Src),
	}

	// 添加导入目录
	for _, include := range p.Includes {
		args = append(args, fmt.Sprintf(`--proto_path=%s`, include))
	}

	// 添加标签
	for _, tag := range p.tags() {
		args = append(args, fmt.Sprintf(`--%s`, tag))
	}

	// 添加插件他输出目录
	args = append(args, fmt.Sprintf(`--%s_out=%s%s`, lang, p.plugins(lang), p.output(lang)))
	// 添加选项
	args = append(args, fmt.Sprintf(`--%s_opt=%s`, lang, p.Opt[lang]))
	// 编译
	err = p.protobuf(args)

	return
}
