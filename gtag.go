package main

import (
	`os`
	`os/exec`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func gtag(conf *config, path string, logger simaqian.Logger) (err error) {
	args := []string{`-input=`, path}
	cmd := exec.Command(`gtag`, `-input=`, path)
	if conf.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	fields := conf.Fields().Connect(field.Strings(`args`, args...))
	if err = cmd.Run(); nil != err {
		logger.Error(`执行处理Golang标签命令出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`执行处理Golang标签命令成功`, fields...)
	}

	return
}
