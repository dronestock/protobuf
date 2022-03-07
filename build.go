package main

import (
	`fmt`
	`strings`
)

func (p *plugin) build(lang string) (err error) {
	args := []interface{}{
		// 加入当前目录
		// 防止出现错误：File does not reside within any path specified using --proto_path
		fmt.Sprintf(`--proto_path=%s`, p.Input),
	}

	// 添加导入目录
	includes := p.includesCache[lang]
	if 0 != len(includes) {
		for _, include := range includes {
			args = append(args, fmt.Sprintf(`--proto_path=%s`, include))
		}
	}

	// 添加标签
	if 0 < len(p.Tags) {
		for _, tag := range p.Tags {
			args = append(args, fmt.Sprintf(`--%s`, tag))
		}
	}

	// 添加插件
	var pb strings.Builder
	plugins := p.pluginsCache[lang]
	if 0 < len(plugins) {
		if langGo == lang || langGogo == lang {
			pb.WriteString(`plugins=`)
		}
		pb.WriteString(strings.Join(plugins, `,`))
		pb.WriteString(`:`)
	}

	// 加入输出目录
	args = append(args, fmt.Sprintf(`--%s_out=%s%s`, lang, pb.String(), p.output(lang)))

	// 添加选项
	opts := p.optsCache[lang]
	if 0 < len(opts) {
		args = append(args, fmt.Sprintf(`--%s_opt=%s`, lang, strings.Join(opts, `,`)))
	}
	err = p.protobuf(args)

	return
}
