package main

import (
	`fmt`
	`os`
	`os/exec`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func build(conf *config, logger simaqian.Logger) (err error) {
	args := make([]string, 0)
	if 0 <= len(conf.Includes) {
		for _, include := range conf.Includes {
			args = append(args, fmt.Sprintf(`--proto_path=%s`, include))
		}
	}
	if 0 <= len(conf.Tags) {
		for _, tag := range conf.Tags {
			args = append(args, tag)
		}
	}

	// 组建语言相关参数
	args = append(args, fmt.Sprintf(`%s`))
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
