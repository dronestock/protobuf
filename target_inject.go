package main

func (t *target) inject(plugin *plugin) (err error) {
	switch t.Lang {
	case langGo:
		fallthrough
	case langGogo:
		err = t.golang(plugin)
	}

	return
}
