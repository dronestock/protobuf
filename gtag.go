package main

import (
	`fmt`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) gtag(path string, logger simaqian.Logger) (err error) {
	args := []string{
		fmt.Sprintf(`-input=%s`, path),
	}
	if p.config.Verbose {
		args = append(args, `-verbose`)
	}

	fields := gox.Fields{
		field.String(`exe`, gtagExe),
		field.String(`path`, path),
		field.Strings(`args`, args...),
	}

	// 记录日志
	logger.Info(`开始处理Golang标签`, fields...)

	options := gex.NewOptions(gex.Args(args...))
	if p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(gtagExe, options...); nil != err {
		logger.Error(`处理Golang标签出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`处理Golang标签完成`, fields...)
	}

	return
}
