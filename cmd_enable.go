package main

import "errors"

type enableCmd struct {
	_    struct{} `help:"change comment -> normal entry if exists"`
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

	dirty := false
	for i, e := range el {
		found := false
		if c.IP != nil && c.Host != nil && e.IP == *c.IP && e.Host == *c.Host {
			found = true
		} else if c.IP != nil && e.IP == *c.IP {
			found = true
		} else if c.Host != nil && e.Host == *c.Host {
			found = true
		}

		if found && e.Type == HostEntry {
			el[i].Enabled = true
			dirty = true
		}
	}

	if dirty {
		return WriteEntries(g.Hosts, el)
	}
	return nil
}
