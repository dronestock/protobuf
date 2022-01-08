package main

import (
	`fmt`
	`os`
	`os/exec`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func gtag(conf *config, path string, logger simaqian.Logger) (err error) {
	args := []string{fmt.Sprintf(`-input=%s`, path), `-verbose`}
	cmd := exec.Command(`gtag`, args...)
	if conf.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	pathField := field.String(`path`, path)
	if err = cmd.Run(); nil != err {
		logger.Error(`处理Golang标签出错`, pathField, field.Strings(`args`, args...), field.Error(err))
	} else {
		logger.Info(`处理Golang标签成功`, pathField)
	}

	return
}
