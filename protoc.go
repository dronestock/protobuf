package main

import (
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

func (p *plugin) protoc(path string, args ...interface{}) (err error) {
	fields := gox.Fields{
		field.String(`exe`, protocExe),
		field.String(`path`, path),
	}
	if err = p.Exec(protocExe, drone.Args(args...), drone.Dir(filepath.Dir(path))); nil != err {
		p.Error(`编译Protobuf文件出错`, fields.Connect(field.Any(`args`, args)).Connect(field.Error(err))...)
	} else {
		p.Info(`编译Protobuf文件完成`, fields...)
	}

	return
}
