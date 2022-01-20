package main

import (
	`github.com/dronestock/drone`
)

const (
	protocExe = `protoc`
	gtagExe   = `gtag`

	protoFilePattern   = `*.proto`
	protoGoFilePattern = `*.pb.go`
	dartLibFilename    = `lib`
)

type plugin struct {
	config *config
}

func newPlugin() drone.Plugin {
	return &plugin{
		config: new(config),
	}
}

func (p *plugin) Configuration() drone.Configuration {
	return p.config
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.builds, drone.Name(`编译Protobuf源文件`)),
		drone.NewStep(p.inject, drone.Name(`编译后源码注入`)),
		drone.NewStep(p.copies, drone.Name(`复制文件`)),
	}
}
