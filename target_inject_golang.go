package main

import (
	"github.com/goexl/gfx"
)

func (t *target) _golang(plugin *plugin) (err error) {
	var filenames []string
	if filenames, err = gfx.All(t._output(), gfx.Pattern(protoGoFilePattern)); nil != err {
		return
	}

	for _, filename := range filenames {
		if err = plugin.gtag(filename); nil != err {
			break
		}
	}

	return
}
