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
