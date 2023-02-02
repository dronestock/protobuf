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

func (t *target) opt(plugin *plugin) (opt []string) {
	if "" == t.Opt {
		return
	}

	opt = make([]string, 0, 1)
	plugins := t.plugins(plugin)
	opt = append(opt, fmt.Sprintf("--%s_opt=%s", t.Lang, t.Opt))
	for _, _plugin := range plugins {
		opt = append(opt, fmt.Sprintf("--%s_opt=%s", _plugin, t.Opt))
	}

	return
}

func (t *target) out(plugin *plugin) (out []string) {
	out = make([]string, 0, 1)
	output := t.output()
	plugins := t.plugins(plugin)
	switch t.Lang {
	case langGo, langGolang, langJava:
		out = append(out, fmt.Sprintf("--%s_out=%s", t.Lang, output))
		for _, _plugin := range plugins {
			out = append(out, fmt.Sprintf("--%s_out=%s", _plugin, output))
		}
	default:
		out = append(out, fmt.Sprintf("--%s_out=%s:%s", t.Lang, strings.Join(plugins, separator), output))
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

func (t *target) plugins(plugin *plugin) (plugins []string) {
	plugins = t.Plugins
	if !*plugin.Defaults {
		return
	}

	switch t.Lang {
	case langJava:
		plugins = append(plugins, "grpc-java", "grpc-gateway")
	case langGo, langGogo:
		plugins = append(plugins, "go-grpc", "grpc-gateway")
	case langDart:
		plugins = append(plugins, "generate_kythe_info")
	case langJs:
		plugins = append(plugins, "binary")
	}
	plugins = append(plugins, plugin.Plugins...)

	return
}
