package main

type lint struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 配置文件
	Config string `default:".protolint.yaml" json:"config"`
}
