package main

import (
	`github.com/storezhang/gfx`
)

func (p *plugin) golang(output string) (err error) {
	var filenames []string
	if filenames, err = gfx.All(output, gfx.Pattern(protoGoFilePattern)); nil != err {
		return
	}

	for _, filename := range filenames {
		if err = p.gtag(filename); nil != err {
			break
		}
	}

	return
}
