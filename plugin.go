package main

import (
	`fmt`
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

	// 输入目录列表
	Inputs []string `default:"${PLUGIN_INPUTS=${INPUTS}}" validate:"required_without=Input"`
	// 输出目录列表
	Outputs []string `default:"${PLUGIN_OUTPUTS=${OUTPUTS}}" validate:"required_without=Output"`

	// 第三方库列表
	Includes []string `default:"${PLUGIN_INCLUDES=${INCLUDES=[]}}"`
	// 标签列表
	Tags []string `default:"${PLUGIN_TAGS=${TAGS}}"`
	// 插件列表
	Plugins []string `default:"${PLUGIN_PLUGINS=${PLUGINS}}"`
	// 选项
	Opts []string `default:"${PLUGIN_OPTS=${OPTS}}"`

	// 额外特性
	// 文件复制列表，在执行完所有操作后，将输入目录的文件或者目录复制到输出目录
	Copies []string `default:"${PLUGIN_COPIES=${COPIES=['README.md', 'LICENSE']}}"`

	inputsCache  map[string][]string
	outputCache  map[string]string
	pluginsCache map[string][]string
	optsCache    map[string][]string
}

func newPlugin() drone.Plugin {
	return &plugin{
		inputsCache:  make(map[string][]string),
		pluginsCache: make(map[string][]string),
		outputCache:  make(map[string]string),
		optsCache:    make(map[string][]string),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	if p.Defaults {
		p.pluginsCache[langGo] = []string{`grpc`}
		p.pluginsCache[langGogo] = []string{`grpc`}
		p.pluginsCache[langDart] = []string{`generate_kythe_info`}
		p.pluginsCache[langJs] = []string{`binary`}

		p.Tags = append(p.Tags, `experimental_allow_proto3_optional`)
	}

	if `` != p.Lang {
		if 0 == len(p.Inputs) {
			p.Inputs = append(p.Inputs, fmt.Sprintf(`%s => %s`, p.Lang, p.Input))
		}
		if 0 == len(p.Outputs) {
			p.Outputs = append(p.Outputs, fmt.Sprintf(`%s => %s`, p.Lang, p.Output))
		}
	}

	// 将原始数据转换成映射
	p.Parses(p.inputsCache, p.Inputs...)
	p.Parses(p.pluginsCache, p.Plugins...)
	p.Parse(p.outputCache, p.Outputs...)
	p.Parses(p.optsCache, p.Opts...)

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

		field.Strings(`inputs`, p.Inputs...),
		field.Strings(`outputs`, p.Outputs...),

		field.Strings(`includes`, p.Includes...),
		field.Strings(`tags`, p.Tags...),
		field.Strings(`plugins`, p.Plugins...),
		field.Strings(`opts`, p.Opts...),

		field.Strings(`copies`, p.Copies...),
	}
}

func (p *plugin) output(lang string) (output string) {
	output = p.outputCache[lang]
	if !p.Defaults {
		return
	}

	switch {
	case langDart == lang && !strings.HasSuffix(output, dartLibFilename):
		output = filepath.Join(output, dartLibFilename)
	case langJava == lang && !strings.HasSuffix(output, javaSourceFilename):
		output = filepath.Join(output, javaSourceFilename)
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
