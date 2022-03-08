package main

import (
	`path/filepath`
	`strings`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 语言
	// nolint:lll
	Lang string `default:"${PLUGIN_LANG=${LANG=go}}" validate:"required_without=Inputs,oneof=go gogo golang java js dart swift python"`
	// 输入目录
	Input string `default:"${PLUGIN_INPUT=${INPUT=.}}"`
	// 输出目录
	Output string `default:"${PLUGIN_OUTPUT=${OUTPUT=.}}"`
	// 输出目录列表
	Outputs map[string]string `default:"${PLUGIN_OUTPUTS=${OUTPUTS}}" validate:"required_without=Output"`

	// 第三方库列表
	Includes []string `default:"${PLUGIN_INCLUDES=${INCLUDES}}"`
	// 标签列表
	Tags []string `default:"${PLUGIN_TAGS=${TAGS}}"`
	// 插件列表
	Plugins map[string][]string `default:"${PLUGIN_PLUGINS=${PLUGINS}}"`
	// 选项
	Opts map[string][]string `default:"${PLUGIN_OPTS=${OPTS}}"`

	// 额外特性
	// 文件复制列表，在执行完所有操作后，将输入目录的文件或者目录复制到输出目录
	Copies []string `default:"${PLUGIN_COPIES=${COPIES=['README.md', 'LICENSE']}}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	if p.Defaults {
		p.Plugins[langGo] = append(p.Plugins[langGo], `grpc`)
		p.Plugins[langGogo] = append(p.Plugins[langGogo], `grpc`)
		p.Plugins[langDart] = append(p.Plugins[langDart], `generate_kythe_info`)
		p.Plugins[langJs] = append(p.Plugins[langJs], `binary`)

		p.Tags = append(p.Tags, `experimental_allow_proto3_optional`)
	}

	if `` != p.Lang && 0 == len(p.Outputs) {
		p.Outputs[p.Lang] = p.Output
	}

	return
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.builds, drone.Name(`编译Protobuf源文件`)),
		drone.NewStep(p.injects, drone.Name(`编译后源码注入`)),
		drone.NewStep(p.copies, drone.Name(`复制文件`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`lang`, p.Lang),
		field.String(`input`, p.Input),
		field.Strings(`output`, p.Output),
		field.Any(`outputs`, p.Outputs),

		field.Strings(`includes`, p.Includes...),
		field.Strings(`tags`, p.Tags...),
		field.Any(`plugins`, p.Plugins),
		field.Any(`opts`, p.Opts),

		field.Strings(`copies`, p.Copies...),
	}
}

func (p *plugin) output(lang string) (output string) {
	output = p.Outputs[lang]
	if !p.Defaults {
		return
	}

	switch {
	case langDart == lang && !strings.HasSuffix(output, dartLibFilename):
		output = filepath.Join(output, dartLibFilename)
	case langJava == lang && !strings.HasSuffix(output, filepath.FromSlash(javaSourceFilename)):
		output = filepath.Join(output, filepath.FromSlash(javaSourceFilename))
	}

	return
}

func (p *plugin) buildable(path string) (buildable bool) {
	buildable = true
	for _, include := range p.Includes {
		if strings.HasPrefix(filepath.Dir(path), include) {
			buildable = false
		}
		if !buildable {
			break
		}
	}

	return
}
