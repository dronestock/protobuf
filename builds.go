package main

func (p *plugin) builds() (undo bool, err error) {
	for lang, inputs := range p.inputsCache {
		for _, input := range inputs {
			if err = p.build(lang, input); nil != err {
				return
			}
		}
	}

	return
}
