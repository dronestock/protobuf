package main

import (
	`fmt`
	`path/filepath`
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
			args = append(args, fmt.Sprintf(`--%s`, tag))
		}
	}

	// 添加插件
	var pluginsBuilder strings.Builder
	plugins := conf.pluginsCache[lang]
	if 0 < len(plugins) {
		if langGo == lang || langGogo == lang {
			pluginsBuilder.WriteString(`plugins=`)
		}
		pluginsBuilder.WriteString(strings.Join(plugins, `,`))
		pluginsBuilder.WriteString(`:`)
	}

	// Dart语言规定，必须打进lib子目录下才能被外部正常引用
	if langDart == lang {
		output = filepath.Join(output, conf.dartLibFilename)
	}
	args = append(args, fmt.Sprintf(`--%s_out=%s%s`, lang, pluginsBuilder.String(), output))

	// 添加选项
	opts := conf.optsCache[lang]
	if 0 < len(opts) {
		args = append(args, fmt.Sprintf(`--%s_opt=%s`, lang, strings.Join(opts, `,`)))
	}
	err = protobuf(conf, input, args, logger)

	return
}
