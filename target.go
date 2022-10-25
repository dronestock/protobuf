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
	Plugins []string `json:"plugins"`
	// 选项
	Opt string `json:"opt"`
}

func (t *target) opt() string {
	return fmt.Sprintf(`--%s_opt=%s`, t.Lang, t.Opt)
}

func (t *target) out(defaults bool) (out []string) {
	out = make([]string, 0, 1)
	prefix, plugins := t.plugins(defaults)
	switch t.Lang {
	case langJava:
		out = append(out, fmt.Sprintf(`--java_out=%s`, t.output()))
		for _, _plugin := range plugins {
			out = append(out, fmt.Sprintf(`--%s-java_out=%s`, fmt.Sprintf(`%s%s`, prefix, _plugin), t.output()))
		}
	default:
		_plugin := fmt.Sprintf(`%s%s`, prefix, strings.Join(plugins, separator))
		out = append(out, fmt.Sprintf(`--%s_out=%s:%s`, t.Lang, _plugin, t.output()))
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

func (t *target) plugins(defaults bool) (prefix string, plugins []string) {
	plugins = t.Plugins
	if !defaults {
		return
	}

	switch t.Lang {
	case langJava:
		plugins = append(plugins, `grpc`)
	case langGo, langGogo:
		prefix = `plugins=`
		plugins = append(plugins, `grpc`)
	case langDart:
		plugins = append(plugins, `generate_kythe_info`)
	case langJs:
		plugins = append(plugins, `binary`)
	}

	return
}
