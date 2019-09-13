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

	var cmp func(string, string) bool
	if c.IP != nil && c.Host != nil {
		cmp = func(ip, host string) bool {
			return ip == *c.IP && host == *c.Host
		}
	} else if c.IP != nil {
		cmp = func(ip, host string) bool {
			return ip == *c.IP
		}
	} else if c.Host != nil {
		cmp = func(ip, host string) bool {
			return host == *c.Host
		}
	}

	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	dirty := false
	for i, e := range el {
		found := false
		if cmp(e.IP, e.Host) {
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
