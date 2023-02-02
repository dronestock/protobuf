package main

import (
	"context"
)

type stepInject struct {
	*plugin
}

func newInjectStep(plugin *plugin) *stepInject {
	return &stepInject{
		plugin: plugin,
	}
}

func (i *stepInject) Runnable() bool {
	return true
}

func (i *stepInject) Run(_ context.Context) (err error) {
	for _, _target := range i.Targets {
		if err = _target.inject(i.plugin); nil != err {
			return
		}
	}

	return
}
