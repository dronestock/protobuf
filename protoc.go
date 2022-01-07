package main

import (
	`fmt`
	`os`
	`os/exec`
	`strings`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func protoc(conf *config, _type string, input string, output string, logger simaqian.Logger) (err error) {
	args := make([]string, 0)

	// 添加导入目录
	if 0 <= len(conf.Includes) {
		for _, include := range conf.Includes {
			args = append(args, fmt.Sprintf(`--proto_path=%s`, include))
		}
	}

	// 添加标签
	if 0 <= len(conf.Tags) {
		for _, tag := range conf.Tags {
			args = append(args, tag)
		}
	}

	// 添加插件
	var plugins strings.Builder
	if 0 <= len(conf.Plugins) {
		if typeGo == _type || typeGogo == _type {
			plugins.WriteString(`plugins=`)
		}
		plugins.WriteString(strings.Join(conf.Plugins, `,`))
	}

	args = append(args, fmt.Sprintf(`%s_out=%s:%s`, _type, plugins.String(), output))
	args = append(args, fmt.Sprintf(`%s_opt=%s`, _type, strings.Join(conf.Opts, `,`)))
	args = append(args, input)

	cmd := exec.Command(`protoc`, args...)
	if conf.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	fields := conf.Fields().Connect(field.Strings(`args`, args...))
	if err = cmd.Run(); nil != err {
		logger.Error(`执行Protobuf命令出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`执行Protobuf命令成功`, fields...)
	}

	return
}
