package main

import (
	`fmt`
	`path/filepath`
	`strings`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type config struct {
	drone.Config

	// 语言
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
	Copies []string `default:"${PLUGIN_COPIES=${COPIES=['README.md']}}"`

	inputsCache  map[string][]string
	outputCache  map[string]string
	pluginsCache map[string][]string
	optsCache    map[string][]string
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`lang`, c.Lang),
		field.String(`input`, c.Input),
		field.Strings(`output`, c.Output),

		field.Strings(`inputs`, c.Inputs...),
		field.Strings(`outputs`, c.Outputs...),

		field.Strings(`includes`, c.Includes...),
		field.Strings(`tags`, c.Tags...),
		field.Strings(`plugins`, c.Plugins...),
		field.Strings(`opts`, c.Opts...),

		field.Strings(`copies`, c.Copies...),
	}
}

func (c *config) Setup() (unset bool, err error) {
	c.init()
	c.Parses(c.inputsCache, c.Inputs...)
	c.Parses(c.pluginsCache, c.Plugins...)
	c.Parse(c.outputCache, c.Outputs...)
	c.Parses(c.optsCache, c.Opts...)

	return
}

func (c *config) output(lang string) (output string) {
	output = c.Output
	if langDart == lang && c.Defaults {
		output = filepath.Join(output, dartLibFilename)
	}

	return
}

func (c *config) buildable(path string) (buildable bool) {
	buildable = true
	for _, include := range c.Includes {
		if strings.HasPrefix(filepath.Dir(path), include) {
			buildable = false
			break
		}
	}

	return
}

func (c *config) init() {
	c.inputsCache = make(map[string][]string)
	c.pluginsCache = make(map[string][]string)
	c.outputCache = make(map[string]string)
	c.optsCache = make(map[string][]string)

	if c.Defaults {
		c.pluginsCache[langGo] = []string{`grpc`}
		c.pluginsCache[langGogo] = []string{`grpc`}
		c.pluginsCache[langDart] = []string{`generate_kythe_info`}
		c.pluginsCache[langJs] = []string{`binary`}

		c.Tags = append(c.Tags, `experimental_allow_proto3_optional`)
	}

	if `` != c.Lang && 0 == len(c.Inputs) {
		c.Inputs = append(c.Inputs, fmt.Sprintf(`%s => %s`, c.Lang, c.Input))
		c.Outputs = append(c.Outputs, fmt.Sprintf(`%s => %s`, c.Lang, c.Output))
	}
}
