package main

import (
	"context"
)

type stepCopy struct {
	*plugin
}

func newCopyStep(plugin *plugin) *stepCopy {
	return &stepCopy{
		plugin: plugin,
	}
}

func (c *stepCopy) Runnable() bool {
	return *c.Defaults || 0 != len(c.Copies)
}

func (c *stepCopy) Run(_ context.Context) (err error) {
	defaults := c.Copies
	if *c.Defaults {
		defaults = append(defaults, "README.md", "LICENSE", "logo.*")
	}

	for _, _target := range c.Targets {
		err = _target.copy(c.Source, c.Logger, defaults...)
		if nil != err {
			return
		}
	}

	return
}
