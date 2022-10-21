package main

func (p *plugin) copy() (undo bool, err error) {
	defaults := p.Copies
	if p.Defaults {
		defaults = append(defaults, "README.md", "LICENSE", "logo.*")
	}

	for _, _target := range p.Targets {
		err = _target.copy(p.Source, p.Logger, defaults...)
		if nil != err {
			return
		}
	}

	return
}
