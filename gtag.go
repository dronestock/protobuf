package main

import (
	`fmt`
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

func (p *plugin) gtag(path string) (err error) {
	args := []interface{}{
		fmt.Sprintf(`-input=%s`, path),
	}
	if p.Verbose {
		args = append(args, `-verbose`)
	}

	fields := gox.Fields{
		field.String(`exe`, gtagExe),
		field.String(`path`, path),
	}
	if err = p.Exec(gtagExe, drone.Args(args...), drone.Dir(filepath.Dir(path))); nil != err {
		p.Error(`处理Golang标签出错`, fields.Connect(field.Any(`args`, args)).Connect(field.Error(err))...)
	} else {
		p.Info(`处理Golang标签完成`, fields...)
	}

	return
}
