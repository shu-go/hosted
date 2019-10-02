package main

import "errors"

type disableCmd struct {
	_    struct{} `help:"change normal entry -> comment if exists"`
	IP   *string
	Host *string
}

func (c disableCmd) Run(g globalCmd) error {
	if c.IP == nil && c.Host == nil {
		return errors.New("IP or Host is required")
	}

	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	dirty := false
	for i, e := range el {
		found := false
		if matches(e, c.IP, c.Host) {
			found = true
		}

		if found && e.Type == HostEntry {
			el[i].Enabled = false
			dirty = true
		}
	}

	if dirty {
		return WriteEntries(g.Hosts, el)
	}
	return nil
}
