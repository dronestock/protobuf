package main

import (
	`fmt`
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

func (p *plugin) gtag(filename string) (err error) {
	args := []interface{}{
		fmt.Sprintf(`-input=%s`, filename),
	}
	if p.Verbose {
		args = append(args, `-verbose`)
	}

	fields := gox.Fields{
		field.String(`exe`, gtagExe),
		field.String(`filename`, filename),
	}
	if err = p.Exec(gtagExe, drone.Args(args...), drone.Dir(filepath.Dir(p.Src))); nil != err {
		p.Error(`注入出错`, fields.Connect(field.Any(`args`, args)).Connect(field.Error(err))...)
	} else if p.Verbose {
		p.Info(`注入成功`, fields...)
	}

	return
}
