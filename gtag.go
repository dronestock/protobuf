package main

import (
	"fmt"
	"path/filepath"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (p *plugin) gtag(filename string) (err error) {
	args := []interface{}{
		fmt.Sprintf("-input=%s", filename),
	}
	if p.Verbose {
		args = append(args, "-verbose")
	}

	fields := gox.Fields[any]{
		field.New("exe", gtagExe),
		field.New("filename", filename),
	}
	if err = p.Command(gtagExe).Args(args...).Dir(filepath.Dir(p.Source)).Exec(); nil != err {
		p.Error("注入出错", fields.Connect(field.New("args", args)).Connect(field.Error(err))...)
	} else if p.Verbose {
		p.Info("注入成功", fields...)
	}

	return
}
