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
	Target target `default:"${PLUGIN_TARGET=${TARGET}}"`
	// 目标列表
	Targets []target `default:"${PLUGIN_TARGETS=${TARGETS}}"`

	// 第三方库列表
	Includes []string `default:"${PLUGIN_INCLUDES=${INCLUDES}}"`
	// 标签列表
	Tags []string `default:"${PLUGIN_TAGS=${TAGS}}"`

	// 额外特性
	// 文件复制列表，在执行完所有操作后，将输入目录的文件或者目录复制到输出目录
	Copies []string `default:"${PLUGIN_COPIES=${COPIES}}"`
}

func newPlugin() drone.Plugin {
	return &plugin{
		Targets: make([]target, 0, 1),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	if 0 == len(p.Targets) {
		p.Targets = append(p.Targets, p.Target)
	}

	return
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewDefaultDelayStep(),
		drone.NewStep(p.builds, drone.Name(`编译`)),
		drone.NewStep(p.injects, drone.Name(`注入`)),
		drone.NewStep(p.copies, drone.Name(`复制`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`input`, p.Source),
		field.Any(`builds`, p.Targets),

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
