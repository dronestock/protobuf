package main

import (
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/goexl/gox`
	`github.com/goexl/gox/field`
)

func (p *plugin) protoc(input string, filename string, args ...interface{}) (err error) {
	fields := gox.Fields{
		field.String(`exe`, protocExe),
		field.String(`filename`, filename),
	}
	// 将需要编译的文件加入到最终的参数中
	args = append(args, filename)

	if err = p.Exec(protocExe, drone.Args(args...), drone.Dir(filepath.Dir(input))); nil != err {
		p.Error(`编译出错`, fields.Connect(field.Any(`args`, args)).Connect(field.Error(err))...)
	} else if p.Verbose {
		p.Info(`编译成功`, fields...)
	}

	return
}
