package main

func (p *plugin) injects() (undo bool, err error) {
	for lang, inputs := range p.inputsCache {
		for _, input := range inputs {
			if err = p.inject(lang, input); nil != err {
				return
			}
		}
	}

	return
}
