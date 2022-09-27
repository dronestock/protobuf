package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type target struct {
	// 语言
	Lang string `default:"go" json:"lang" validate:"oneof=go gogo golang java js dart swift python"`
	// 输出目录
	Output string `default:"." json:"output"`
	// 插件列表
	Plugins string `json:"plugins"`
	// 选项
	Opt string `json:"opt"`
}

func (t *target) opt() string {
	return fmt.Sprintf(`--%s_opt=%s`, t.Lang, t.Opt)
}

func (t *target) out(defaults bool) string {
	return fmt.Sprintf(`--%s_out=%s%s`, t.Lang, t._plugins(defaults), t._output())
}

func (t *target) _output() (output string) {
	output = t.Output

	switch {
	case langDart == t.Lang && !strings.HasSuffix(output, dartLibFilename):
		output = filepath.Join(output, dartLibFilename)
	case typeJava == t.Lang && !strings.HasSuffix(output, filepath.FromSlash(javaSourceFilename)):
		output = filepath.Join(output, filepath.FromSlash(javaSourceFilename))
	}

	// 转换成绝对路径，防止Protobuf找不到路径报错
	output, _ = filepath.Abs(output)
	_ = os.MkdirAll(output, os.ModePerm)

	return
}

func (t *target) _plugins(defaults bool) (plugins string) {
	plugins = t.Plugins
	if !defaults {
		return
	}

	var dps string
	prefix := ``
	switch t.Lang {
	case langGo, langGogo:
		dps = `grpc`
		prefix = `plugins=`
	case langDart:
		dps = `generate_kythe_info`
	case langJs:
		dps = `binary`
	default:
		return
	}

	olds := make([]string, 0)
	if `` != strings.TrimSpace(plugins) {
		olds = append(olds, strings.Split(plugins, separator)...)
	}
	if `` != dps && !strings.Contains(plugins, dps) {
		olds = append(olds, dps)
	}
	plugins = fmt.Sprintf(`%s%s:`, prefix, strings.Join(olds, separator))

	return
}
