package main

type descriptor struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 要合并的源文件
	Source string `json:"source"`
	// 要合并的源文件列表
	Sources []string `json:"sources" validate:"required_if=Source"`
	// 输出文件
	Output string `default:"descriptor.pb" json:"output"`
	// 选项
	Opts []string `json:"OPTS"`
}

func (p *plugin) descriptor() (undo bool, err error) {
	if undo = !*p.Descriptor.Enabled || 0 == len(p.Descriptors); undo {
		return
	}

	for _, _descriptor := range p.Descriptors {
		if err = _descriptor.build(p); nil != err {
			return
		}
	}

	return
}
