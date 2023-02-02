package main

import (
	"context"
)

type stepDescriptor struct {
	*plugin
}

func newDescriptorStep(plugin *plugin) *stepDescriptor {
	return &stepDescriptor{
		plugin: plugin,
	}
}

func (d *stepDescriptor) Runnable() bool {
	return 0!=len(d.Descriptors)
}

func (d *stepDescriptor) Run(_ context.Context) (err error) {
	for _, _descriptor := range d.Descriptors {
		if _descriptor.enabled() {
			err = _descriptor.build(d.plugin)
		}

		if nil != err {
			return
		}
	}

	return
}
