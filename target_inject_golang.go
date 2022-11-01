package main

import (
	"github.com/goexl/gfx"
)

func (t *target) golang(plugin *plugin) (err error) {
	var filenames []string
	if filenames, err = gfx.All(t.output(), gfx.Suffix(protoGoFileSuffix)); nil != err {
		return
	}

	for _, filename := range filenames {
		if err = plugin.gtag(filename); nil != err {
			break
		}
	}

	return
}
