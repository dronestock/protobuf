package main

import (
	`github.com/storezhang/gfx`
)

func (p *plugin) golang(output string) (err error) {
	var files []string
	if files, err = gfx.All(output, gfx.Pattern(protoGoFilePattern)); nil != err {
		return
	}

	for _, file := range files {
		if err = p.gtag(file); nil != err {
			break
		}
	}

	return
}
