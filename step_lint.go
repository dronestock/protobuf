package main

import (
	"context"
	"path/filepath"

	"github.com/goexl/gfx"
)

type stepLint struct {
	*plugin
}

func newLintStep(plugin *plugin) *stepLint {
	return &stepLint{
		plugin: plugin,
	}
}

func (l *stepLint) Runnable() bool {
	return !(nil != l.Lint.Enabled && !*l.Lint.Enabled)
}

func (l *stepLint) Run(_ context.Context) (err error) {
	config := lintConfigFilename
	if final, exists := gfx.Exists(filepath.Join(l.Source, l.Lint.Config)); exists {
		config = final
	}
	err = l.Command(lintExe).Args("lint", "-fix", "-config_path", config, l.Source).Exec()

	return
}
