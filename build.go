package main

import (
	`fmt`
	`strings`

	`github.com/storezhang/simaqian`
)

func (p *plugin) build(lang string, input string, output string, logger simaqian.Logger) (err error) {
	args := []string{
		// 加入当前目录
		// 防止出现File does not reside within any path specified using --proto_path的错误
		fmt.Sprintf(`--proto_path=%s`, input),
	}

	// 添加导入目录
	if 0 < len(p.config.Includes) {
		for _, include := range p.config.Includes {
			args = append(args, fmt.Sprintf(`--proto_path=%s`, include))
		}
	}

	// 添加标签
	if 0 < len(p.config.Tags) {
		for _, tag := range p.config.Tags {
			args = append(args, fmt.Sprintf(`--%s`, tag))
		}
	}

	// 添加插件
	var pluginsBuilder strings.Builder
	plugins := p.config.pluginsCache[lang]
	if 0 < len(plugins) {
		if langGo == lang || langGogo == lang {
			pluginsBuilder.WriteString(`plugins=`)
		}
		pluginsBuilder.WriteString(strings.Join(plugins, `,`))
		pluginsBuilder.WriteString(`:`)
	}

	// 加入输出目录
	args = append(args, fmt.Sprintf(`--%s_out=%s%s`, lang, pluginsBuilder.String(), p.config.output(lang)))

	// 添加选项
	opts := p.config.optsCache[lang]
	if 0 < len(opts) {
		args = append(args, fmt.Sprintf(`--%s_opt=%s`, lang, strings.Join(opts, `,`)))
	}
	err = p.protobuf(input, args, logger)

	return
}
