package main

import (
	"path/filepath"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (p *plugin) protoc(input string, filenames []string, args []any) (err error) {
	fields := gox.Fields[any]{
		field.New("exe", protocExe),
		field.New("input", input),
		field.New("filenames", filenames),
	}
	// 有警告时不允许编译通过
	if p.FatalWarnings {
		args = append(args, "--fatal_warnings")
	}

	// 将需要编译的文件加入到最终的参数中
	for _, filename := range filenames {
		if abs, ae := filepath.Abs(filename); nil != ae {
			err = ae
		} else {
			args = append(args, abs)
		}

		if nil != err {
			return
		}
	}

	if err = p.Exec(protocExe, drone.Args(args...), drone.Dir(filepath.Dir(input))); nil != err {
		p.Error("编译出错", fields.Connect(field.New("args", args)).Connect(field.Error(err))...)
	} else if p.Verbose {
		p.Info("编译成功", fields...)
	}

	return
}
