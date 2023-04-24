package main

import (
	"path/filepath"

	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

func (p *plugin) gtag(filename string) (err error) {
	ga := args.New().Long(strike).Build()
	ga.Arg("input", filename)
	if p.Verbose {
		ga.Flag("verbose")
	}

	arguments := ga.Build()
	fields := gox.Fields[any]{
		field.New("exe", p.Binary.Gtag),
		field.New("filename", filename),
	}
	if _, err = p.Command(p.Binary.Gtag).Args(arguments).Dir(filepath.Dir(p.Source)).Build().Exec(); nil != err {
		p.Error("注入出错", fields.Add(field.New("args", arguments.String())).Add(field.Error(err))...)
	} else if p.Verbose {
		p.Info("注入成功", fields...)
	}

	return
}
