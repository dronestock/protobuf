package main

import (
	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) protoc(path string, logger simaqian.Logger, args ...string) (err error) {
	fields := gox.Fields{
		field.String(`exe`, protocExe),
		field.String(`path`, path),
		field.Strings(`args`, args...),
	}

	// 记录日志
	logger.Info(`开始编译Protobuf文件`, fields...)

	options := gex.NewOptions(gex.Args(args...))
	if p.config.Debug {
		options = append(options, gex.Quiet())
	}

	if _, err = gex.Run(protocExe, options...); nil != err {
		logger.Error(`编译Protobuf文件出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`编译Protobuf文件完成`, fields...)
	}

	return
}
