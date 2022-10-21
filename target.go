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

func (t *target) out(defaults bool) (out []string) {
	switch t.Lang {
	case langJava:
		out = []string{
			fmt.Sprintf(`--java_out=%s`, t.output()),
			fmt.Sprintf(`--%sjava_out=%s`, t.plugins(defaults), t.output()),
		}
	default:
		out = []string{fmt.Sprintf(`--%s_out=%s:%s`, t.Lang, t.plugins(defaults), t.output())}
	}

	return
}

func (t *target) output() (output string) {
	output = t.Output

	switch {
	case langDart == t.Lang && !strings.HasSuffix(output, dartLibFilename):
		output = filepath.Join(output, dartLibFilename)
	case langJava == t.Lang && !strings.HasSuffix(output, filepath.FromSlash(javaSourceFilename)):
		output = filepath.Join(output, filepath.FromSlash(javaSourceFilename))
	}

	// 转换成绝对路径，防止Protobuf找不到路径报错
	output, _ = filepath.Abs(output)
	_ = os.MkdirAll(output, os.ModePerm)

	return
}

func (t *target) plugins(defaults bool) (plugins string) {
	plugins = t.Plugins
	if !defaults {
		return
	}

	var dps string
	prefix := ``
	staff := ``
	switch t.Lang {
	case langJava:
		dps = `grpc`
		staff = `-`
	case langGo, langGogo:
		prefix = `plugins=`
		dps = `grpc`
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
	plugins = fmt.Sprintf(`%s%s%s`, prefix, strings.Join(olds, separator), staff)

	return
}
