package main

import (
	"context"
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
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

	la := args.New().Long(strike).Build()
	la.Subcommand("lint")
	la.Flag("fix")
	la.Option("config_path", config)
	la.Add(l.Source)
	_, err = l.Command(l.Binary.Lint).Args(la.Build()).Build().Exec()

	return
}
