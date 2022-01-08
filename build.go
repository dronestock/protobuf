package main

import (
	`fmt`
	`strings`

	`github.com/storezhang/simaqian`
)

func build(conf *config, lang string, input string, output string, logger simaqian.Logger) (err error) {
	args := make([]string, 0)

	// 加入当前目录
	// 防止出现File does not reside within any path specified using --proto_path的错误
	args = append(args, fmt.Sprintf(`--proto_path=%s`, input))
	// 添加导入目录
	if 0 < len(conf.Includes) {
		for _, include := range conf.Includes {
			args = append(args, fmt.Sprintf(`--proto_path=%s`, include))
		}
	}

	// 添加标签
	if 0 < len(conf.Tags) {
		for _, tag := range conf.Tags {
			args = append(args, tag)
		}
	}

	// 添加插件
	var plugins strings.Builder
	if 0 < len(conf.Plugins) {
		if langGo == lang || langGogo == lang {
			plugins.WriteString(`plugins=`)
		}
		plugins.WriteString(`:`)
		plugins.WriteString(strings.Join(conf.Plugins, `,`))
	}

	args = append(args, fmt.Sprintf(`--%s_out=%s%s`, lang, plugins.String(), output))
	if 0 < len(conf.Opts) {
		args = append(args, fmt.Sprintf(`--%s_opt=%s`, lang, strings.Join(conf.Opts, `,`)))
	}
	err = protobuf(conf, input, args, logger)

	return
}
