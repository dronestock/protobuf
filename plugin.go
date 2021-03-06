package main

import (
	`fmt`
	`path/filepath`
	`strings`

	`github.com/dronestock/drone`
	`github.com/goexl/gox`
	`github.com/goexl/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 类型
	// nolint:lll
	Type string `default:"${PLUGIN_TYPE=${TYPE=go}}" validate:"required_without=Inputs,oneof=go gogo golang java js dart swift python"`
	// 源文件目录
	Source string `default:"${PLUGIN_SOURCE=${SOURCE=.}}"`
	// 输出目录
	Output string `default:"${PLUGIN_OUTPUT=${OUTPUT=.}}"`
	// 输出目录列表
	Outputs map[string]string `default:"${PLUGIN_OUTPUTS=${OUTPUTS}}" validate:"required_without=Output"`

	// 第三方库列表
	Includes []string `default:"${PLUGIN_INCLUDES=${INCLUDES}}"`
	// 标签列表
	Tags []string `default:"${PLUGIN_TAGS=${TAGS}}"`
	// 插件列表
	Plugins map[string]string `default:"${PLUGIN_PLUGINS=${PLUGINS}}"`
	// 选项
	Opt map[string]string `default:"${PLUGIN_OPT=${OPT}}"`

	// 额外特性
	// 文件复制列表，在执行完所有操作后，将输入目录的文件或者目录复制到输出目录
	Copies []string `default:"${PLUGIN_COPIES=${COPIES}}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	if 0 == len(p.Outputs) {
		p.Outputs[p.Type] = p.Output
	}

	return
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.builds, drone.Name(`编译`)),
		drone.NewStep(p.injects, drone.Name(`注入`)),
		drone.NewStep(p.copies, drone.Name(`复制`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`type`, p.Type),
		field.String(`input`, p.Source),
		field.Strings(`output`, p.Output),
		field.Any(`outputs`, p.Outputs),

		field.Strings(`includes`, p.Includes...),
		field.Strings(`tags`, p.Tags...),
		field.Any(`plugins`, p.Plugins),
		field.Any(`opt`, p.Opt),

		field.Strings(`copies`, p.Copies...),
	}
}

func (p *plugin) plugins(typ string) (plugins string) {
	plugins = p.Plugins[typ]
	if !p.Defaults {
		return
	}

	var defaults string
	prefix := ``
	switch typ {
	case typeGo, typeGogo:
		defaults = `grpc`
		prefix = `plugins=`
	case typeDart:
		defaults = `generate_kythe_info`
	case typeJs:
		defaults = `binary`
	default:
		return
	}

	olds := make([]string, 0)
	if `` != strings.TrimSpace(plugins) {
		olds = append(olds, strings.Split(plugins, separator)...)
	}
	if `` != defaults && !strings.Contains(plugins, defaults) {
		olds = append(olds, defaults)
	}
	plugins = fmt.Sprintf(`%s%s:`, prefix, strings.Join(olds, separator))

	return
}

func (p *plugin) tags() (tags []string) {
	tags = p.Tags
	if p.Defaults {
		tags = append(tags, `experimental_allow_proto3_optional`)
	}

	return
}

func (p *plugin) _copies() (copies []string) {
	copies = p.Copies
	if p.Defaults {
		copies = append(copies, "README.md", "LICENSE", "logo.*")
	}

	return
}

func (p *plugin) output(typ string) (output string) {
	output = p.Outputs[typ]

	switch {
	case typeDart == typ && !strings.HasSuffix(output, dartLibFilename):
		output = filepath.Join(output, dartLibFilename)
	case typeJava == typ && !strings.HasSuffix(output, filepath.FromSlash(javaSourceFilename)):
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
