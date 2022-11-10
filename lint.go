package main

import (
	"path/filepath"

	"github.com/dronestock/drone"
	"github.com/goexl/gfx"
)

type lint struct {
	// 是否开启
	Enabled bool `default:"true" json:"enabled"`
	// 配置文件
	Config string `default:".protolint.yaml" json:"config"`
}

func (p *plugin) lint() (undo bool, err error) {
	if undo = !p.Lint.Enabled; undo {
		return
	}

	config := lintConfigFilename
	if final, exists := gfx.Exists(filepath.Join(p.Source, p.Lint.Config)); exists {
		config = final
	}
	err = p.Exec(lintExe, drone.Args(`lint`, `-fix`, `-config_path`, config, p.Source))

	return
}
