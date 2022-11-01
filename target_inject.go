package main

func (t *target) inject(plugin *plugin) (err error) {
	switch t.Lang {
	case langGo, langGogo:
		err = t.golang(plugin)
	}

	return
}
