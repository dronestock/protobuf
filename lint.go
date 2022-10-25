package main

import (
	"github.com/dronestock/drone"
)

func (p *plugin) lint() (undo bool, err error) {
	err = p.Exec(protolintExe, drone.Args(`lint`, `-fix`, `-auto_disable=next`, p.Source))

	return
}
