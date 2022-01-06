package main

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/mengpo`
	`github.com/storezhang/validatorx`
)

type config struct {
	// 语言
	Lang string `default:"${PLUGIN_LANG=${LANG=go}}" validate:"required,oneof=go golang java javascript dart"`
	// 目录列表
	Folders []string `default:"${PLUGIN_FOLDERS=${FOLDERS=.}}" validate:"required"`
	// 第三方库列表
	Includes []string `default:"${PLUGIN_INCLUDES=${INCLUDES}}"`
	// 标签列表
	Tags []string `default:"${PLUGIN_TAGS=${TAGS}}"`
	// 插件列表
	Plugins []string `default:"${PLUGIN_PLUGINS=${PLUGINS}}"`
	// 选项
	Opts []string `default:"${PLUGIN_OPTS=${OPTS}}"`

	// 是否启用默认优化
	Defaults bool `default:"${PLUGIN_DEFAULTS=${DEFAULTS=true}}"`
	// 是否显示调试信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`

	defaultOpts    map[lang][]string
	defaultPlugins map[lang][]string
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`lang`, string(c.Lang)),
		field.Strings(`folders`, c.Folders...),
		field.Strings(`includes`, c.Includes...),
		field.Strings(`tags`, c.Tags...),
		field.Strings(`plugins`, c.Plugins...),
		field.Strings(`opts`, c.Opts...),

		field.Bool(`verbose`, c.Verbose),
	}
}

func (c *config) load() (err error) {
	// 处理环境变量为字符串的时候和默认值格式不兼容
	if err = parseEnvs(`FOLDERS`, `INCLUDES`, `TAGS`, `PLUGINS`, `OPTS`); nil != err {
		return
	}
	if err = mengpo.Set(c); nil != err {
		return
	}
	if err = validatorx.Struct(c); nil != err {
		return
	}
	c.defaultOpts = map[lang][]string{
		langGo: {},
	}

	return
}
