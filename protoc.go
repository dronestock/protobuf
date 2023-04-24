package main

import (
	"path/filepath"

	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

func (p *plugin) protoc(input string, filenames []string, args *args.Builder) (err error) {
	fields := gox.Fields[any]{
		field.New("exe", p.Binary.Protoc),
		field.New("input", input),
		field.New("filenames", filenames),
	}
	// 有警告时不允许编译通过
	if p.FatalWarnings {
		args.Flag("fatal_warnings")
	}

	// 将需要编译的文件加入到最终的参数中
	for _, filename := range filenames {
		args.Add(filename)
	}

	arguments := args.Build()
	if _, err = p.Command(p.Binary.Protoc).Args(args.Build()).Dir(filepath.Dir(input)).Build().Exec(); nil != err {
		p.Error("编译出错", fields.Add(field.New("args", arguments.String())).Add(field.Error(err))...)
	} else if p.Verbose {
		p.Info("编译成功", fields...)
	}

	return
}
