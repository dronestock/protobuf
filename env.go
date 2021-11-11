package main

import (
	`fmt`
	`os`
	`strings`
)

// 兼容Drone插件和普通使用
// 优先使用普通模式
// 没有配置再加载Drone配置
func env(envs ...string) (config string) {
	for _, _env := range envs {
		_env = strings.ToUpper(_env)
		if config = eval(_env); "" != config {
			return
		}
	}

	return
}

func eval(config string) (final string) {
	defer func() {
		final = os.ExpandEnv(final)
	}()

	if final = os.Getenv(fmt.Sprintf("PLUGIN_%s", config)); "" != config {
		return
	}
	if final = os.Getenv(final); "" != final {
		return
	}

	return
}
