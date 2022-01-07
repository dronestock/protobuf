package main

import (
	`fmt`
	`strings`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/mengpo`
	`github.com/storezhang/validatorx`
)

type config struct {
	// 类型
	Type string `default:"${PLUGIN_TYPE=${TYPE=go}}" validate:"required_without=Types,oneof=go golang java javascript dart"`
	// 输入目录
	Input string `default:"${PLUGIN_INPUT=${INPUT=.}}"`
	// 输出目录
	Output string `default:"${PLUGIN_OUTPUT=${OUTPUT=.}}"`

	// 类型列表
	Types []string `default:"${PLUGIN_TYPES=${TYPES}}" validate:"required_without=Type,dive,oneof=go golang java javascript dart"`
	// 输入目录列表
	Inputs []string `default:"${PLUGIN_INPUTS=${INPUTS}}" validate:"required_without=Input"`
	// 输出目录列表
	Outputs []string `default:"${PLUGIN_OUTPUTS=${OUTPUTS}}" validate:"required_without=Output"`

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

	inputsCache    map[string][]string
	outputCache    map[string]string
	defaultOpts    map[string][]string
	defaultPlugins map[string][]string
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`type`, c.Type),
		field.String(`input`, c.Input),
		field.Strings(`output`, c.Output),

		field.Strings(`types`, c.Types...),
		field.Strings(`inputs`, c.Inputs...),
		field.Strings(`outputs`, c.Outputs...),

		field.Strings(`includes`, c.Includes...),
		field.Strings(`tags`, c.Tags...),
		field.Strings(`plugins`, c.Plugins...),
		field.Strings(`opts`, c.Opts...),

		field.Bool(`defaults`, c.Defaults),
		field.Bool(`verbose`, c.Verbose),
	}
}

func (c *config) load() (err error) {
	// 处理环境变量为字符串的时候和默认值格式不兼容
	if err = parseEnvs(`INPUTS`, `OUTPUTS`, `INCLUDES`, `TAGS`, `PLUGINS`, `OPTS`); nil != err {
		return
	}
	if err = mengpo.Set(c); nil != err {
		return
	}
	if err = validatorx.Struct(c); nil != err {
		return
	}
	if `` != c.Type {
		c.Types = append(c.Types, c.Type)
		c.Inputs = append(c.Inputs, fmt.Sprintf(`%s => %s`, c.Type, c.Input))
		c.Outputs = append(c.Outputs, fmt.Sprintf(`%s => %s`, c.Type, c.Output))
	}
	for _, input := range c.Inputs {
		c.parseConfig(input, c.putInputs)
	}
	for _, output := range c.Outputs {
		c.parseConfig(output, c.putOutput)
	}

	// 处理默认配置
	c.defaultOpts = map[string][]string{
		`go`: {},
	}

	return
}

func (c *config) parseConfig(original string, put func(configs []string)) {
	var _configs []string
	defer func() {
		put(_configs)
	}()

	if _configs = strings.Split(original, "@"); 2 <= len(_configs) {
		return
	}
	if _configs = strings.Split(original, "=>"); 2 <= len(_configs) {
		return
	}
	if _configs = strings.Split(original, "->"); 2 <= len(_configs) {
		return
	}
	if _configs = strings.Split(original, " "); 2 <= len(_configs) {
		return
	}

	return
}

func (c *config) putInputs(configs []string) {
	if nil != configs && 2 <= len(configs) {
		c.inputsCache[configs[0]] = c.splits(configs[1], `,`, `|`, `||`)
	}

	return
}

func (c *config) putOutput(configs []string) {
	if nil != configs && 2 <= len(configs) {
		c.outputCache[configs[0]] = configs[1]
	}

	return
}

func (c *config) splits(config string, seps ...string) (configs []string) {
	for _, sep := range seps {
		if strings.Contains(config, sep) {
			configs = strings.Split(config, sep)
			break
		}
	}

	return
}
