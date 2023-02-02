package main

import (
	"context"
)

type stepBuild struct {
	*plugin
}

func newBuildStep(plugin *plugin) *stepBuild {
	return &stepBuild{
		plugin: plugin,
	}
}

func (b *stepBuild) Runnable() bool {
	return true
}

func (b *stepBuild) Run(_ context.Context) (err error) {
	for _, _target := range b.Targets {
		if err = _target.build(b.plugin); nil != err {
			return
		}
	}

	return
}
