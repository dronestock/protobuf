package main

import (
	`os`
	`os/exec`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func protoc(conf *config, path string, logger simaqian.Logger, args ...string) (err error) {
	args = append(args, path)
	cmd := exec.Command(`protoc`, args...)
	if conf.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	pathField := field.String(`path`, path)
	if err = cmd.Run(); nil != err {
		logger.Error(`编译Proto文件出错`, pathField, field.Strings(`args`, args...), field.Error(err))
	} else {
		logger.Info(`编译Proto文件成功`, pathField)
	}

	return
}
