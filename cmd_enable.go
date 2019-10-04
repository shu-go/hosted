package main

import (
	"errors"
)

type enableCmd struct {
	_ struct{} `help:"change comment -> normal entry if exists"`

	IP   *string
	Host *string
}

func (c enableCmd) Run(g globalCmd) error {
	if c.IP == nil && c.Host == nil {
		return errors.New("IP or Host is required")
	}

	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	el = c.Enable(el)

	if el != nil {
		return WriteEntries(g.Hosts, el)
	}
	return nil
}

func (c enableCmd) Enable(el []Entry) []Entry {
	dirty := false

	for i, e := range el {
		found := false
		if matches(e, c.IP, c.Host) {
			found = true
		}

		if found && e.Type == HostEntry && !el[i].Enabled {
			el[i].Enabled = true
			dirty = true
		}
	}

	if dirty {
		return el
	}
	return nil
}
