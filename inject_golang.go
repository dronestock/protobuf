package main

import (
	`github.com/storezhang/gfx`
	`github.com/storezhang/simaqian`
)

func (p *plugin) golang(output string, logger simaqian.Logger) (err error) {
	var files []string
	if files, err = gfx.All(output, gfx.Pattern(protoGoFilePattern)); nil != err {
		return
	}

	for _, file := range files {
		if err = p.gtag(file, logger); nil != err {
			break
		}
	}

	return
}
