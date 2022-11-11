package main

import (
	"path/filepath"
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 源文件目录
	Source string `default:"${PLUGIN_SOURCE=${SOURCE=.}}"`
	// 目标
	Target *target `default:"${PLUGIN_TARGET=${TARGET}}" validate:"required_without=Targets"`
	// 目标列表
	Targets []*target `default:"${PLUGIN_TARGETS=${TARGETS}}" validate:"required_without=Target"`

	// 第三方库列表
	Includes []string `default:"${PLUGIN_INCLUDES=${INCLUDES}}"`
	// 标签列表
	Tags []string `default:"${PLUGIN_TAGS=${TAGS}}"`
	// 有警告时不允许编译通过
	FatalWarnings bool `default:"${PLUGIN_FATAL_WARNINGS=${FATAL_WARNINGS=true}}"`

	// 额外特性
	// 静态检查
	Lint lint `default:"${PLUGIN_LINT=${LINT}}"`
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
	if nil != p.Target {
		if nil == p.Targets {
			p.Targets = make([]*target, 0, 1)
		}
		p.Targets = append(p.Targets, p.Target)
	}

	return
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		// 纯静态检查，不需要重试
		drone.NewStep(p.lint, drone.Name(`检查`), drone.Interrupt()),
		// 编译，不依赖网络环境，不需要重试
		drone.NewStep(p.build, drone.Name(`编译`), drone.Interrupt()),
		// 注入，不依赖网络环境，不需要重试
		drone.NewStep(p.inject, drone.Name(`注入`), drone.Interrupt()),
		// 复制，不依赖网络环境，不需要重试
		drone.NewStep(p.copy, drone.Name(`复制`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`input`, p.Source),
		field.Any(`targets`, p.Targets),

		field.Strings(`includes`, p.Includes...),
		field.Strings(`tags`, p.Tags...),

		field.Strings(`copies`, p.Copies...),
	}
}

func (p *plugin) tags() (tags []string) {
	tags = p.Tags
	if p.Defaults {
		tags = append(tags, `experimental_allow_proto3_optional`)
	}

	return
}

func (p *plugin) buildable(path string) (buildable bool, err error) {
	if buildable, err = filepath.Match(protoFilePattern, filepath.Base(path)); nil != err || !buildable {
		return
	}

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
