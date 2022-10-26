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

func (t *target) opt(defaults bool) (opt []string) {
	opt = make([]string, 0, 1)
	plugins := t.plugins(defaults)
	switch t.Lang {
	case langGo, langGolang:
		opt = append(opt, fmt.Sprintf(`--go_opt=%s`, t.Opt))
		for _, _plugin := range plugins {
			opt = append(opt, fmt.Sprintf(`--go-%s_opt=%s`, _plugin, t.Opt))
		}
	default:
		opt = append(opt, fmt.Sprintf(`--%s_opt=%s`, t.Lang, t.Opt))
	}

	return
}

func (t *target) out(defaults bool) (out []string) {
	out = make([]string, 0, 1)
	plugins := t.plugins(defaults)
	switch t.Lang {
	case langJava:
		out = append(out, fmt.Sprintf(`--java_out=%s`, t.output()))
		for _, _plugin := range plugins {
			out = append(out, fmt.Sprintf(`--%s-java_out=%s`, _plugin, t.output()))
		}
	case langGo, langGolang:
		out = append(out, fmt.Sprintf(`--go_out=%s`, t.output()))
		for _, _plugin := range plugins {
			out = append(out, fmt.Sprintf(`--go-%s_out=%s`, _plugin, t.output()))
		}
	default:
		out = append(out, fmt.Sprintf(`--%s_out=%s:%s`, t.Lang, strings.Join(plugins, separator), t.output()))
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

func (t *target) plugins(defaults bool) (plugins []string) {
	plugins = t.Plugins
	if !defaults {
		return
	}

	switch t.Lang {
	case langJava:
		plugins = append(plugins, `grpc`)
	case langGo, langGogo:
		plugins = append(plugins, `grpc`)
	case langDart:
		plugins = append(plugins, `generate_kythe_info`)
	case langJs:
		plugins = append(plugins, `binary`)
	}

	return
}
