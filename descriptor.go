package main

type descriptor struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 要合并的源文件
	Source string `json:"source" validate:"omitempty,required_without=Sources"`
	// 要合并的源文件列表
	Sources []string `json:"sources" validate:"omitempty,required_without=Source"`
	// 输出文件
	Output string `default:"descriptor.pb" json:"output"`
	// 选项
	Opts []string `json:"OPTS"`
}

func (d *descriptor) enabled() bool {
	return nil != d.Enabled && *d.Enabled
}

func (p *plugin) descriptor() (undo bool, err error) {
	if undo = 0 == len(p.Descriptors); undo {
		return
	}

	for _, _descriptor := range p.Descriptors {
		if _descriptor.enabled() {
			err = _descriptor.build(p)
		}

		if nil != err {
			return
		}
	}

	return
}
