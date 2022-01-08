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
	// 语言
	Lang string `default:"${PLUGIN_LANG=${LANG=go}}" validate:"required_without=Inputs,oneof=go golang java javascript dart"`
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

	// 是否启用默认优化
	Defaults bool `default:"${PLUGIN_DEFAULTS=${DEFAULTS=true}}"`
	// 是否显示调试信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`

	inputsCache    map[string][]string
	outputCache    map[string]string
	defaultOpts    map[string][]string
	defaultPlugins map[string][]string

	protoFilePattern   string
	protoGoFilePattern string
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

	// 处理默认配置
	c.inputsCache = make(map[string][]string)
	c.outputCache = make(map[string]string)
	c.defaultOpts = map[string][]string{
		`go`: {},
	}
	c.protoFilePattern = `*.proto`
	c.protoGoFilePattern = `*.pb.go`

	// 解析参数
	if `` != c.Lang && 0 == len(c.Inputs) {
		c.Inputs = append(c.Inputs, fmt.Sprintf(`%s => %s`, c.Lang, c.Input))
		c.Outputs = append(c.Outputs, fmt.Sprintf(`%s => %s`, c.Lang, c.Output))
	}
	for _, input := range c.Inputs {
		c.parseConfig(input, c.putInputs)
	}
	for _, output := range c.Outputs {
		c.parseConfig(output, c.putOutput)
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
		value := strings.TrimSpace(configs[1])
		if `` == value {
			return
		}

		c.inputsCache[strings.TrimSpace(configs[0])] = c.splits(value, `,`, `|`, `||`)
	}

	return
}

func (c *config) putOutput(configs []string) {
	if nil != configs && 2 <= len(configs) {
		value := strings.TrimSpace(configs[1])
		if `` == value {
			return
		}

		c.outputCache[strings.TrimSpace(configs[0])] = value
	}

	return
}

func (c *config) splits(config string, seps ...string) (configs []string) {
	configs = []string{config}
	for _, sep := range seps {
		if strings.Contains(config, sep) {
			configs = strings.Split(config, sep)
			break
		}
	}

	return
}
