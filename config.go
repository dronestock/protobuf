package main

import (
	`fmt`
	`path/filepath`
	`strings`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/mengpo`
	`github.com/storezhang/validatorx`
)

type config struct {
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

	// 是否启用默认优化
	Defaults bool `default:"${PLUGIN_DEFAULTS=${DEFAULTS=true}}"`
	// 是否显示调试信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`

	inputsCache        map[string][]string
	outputCache        map[string]string
	pluginsCache       map[string][]string
	optsCache          map[string][]string

	protoFilePattern   string
	protoGoFilePattern string
	dartLibFilename    string
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
	c.parse()

	return
}

func (c *config) parse() {
	c.init()

	for _, input := range c.Inputs {
		c.parseConfig(input, c.puts(c.inputsCache))
	}
	for _, plugin := range c.Plugins {
		c.parseConfig(plugin, c.puts(c.pluginsCache))
	}
	for _, output := range c.Outputs {
		c.parseConfig(output, c.put(c.outputCache))
	}
	for _, opt := range c.Opts {
		c.parseConfig(opt, c.puts(c.optsCache))
	}
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

func (c *config) puts(cache map[string][]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			cache[strings.TrimSpace(configs[0])] = c.splits(value, `,`, `|`, `||`)
		}
	}
}

func (c *config) put(cache map[string]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			cache[strings.TrimSpace(configs[0])] = value
		}

		return
	}
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

func (c *config) init() {
	c.inputsCache = make(map[string][]string)
	c.pluginsCache = make(map[string][]string)
	c.outputCache = make(map[string]string)
	c.optsCache = make(map[string][]string)

	c.protoFilePattern = `*.proto`
	c.protoGoFilePattern = `*.pb.go`
	c.dartLibFilename = `lib`

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
