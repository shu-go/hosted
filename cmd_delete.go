package main

import (
	"errors"
)

type deleteCmd struct {
	_    struct{} `help:"delete an entry with  --ip 192.168.1.200 and/or --host oldserver"`
	IP   *string
	Host *string
}

func (c deleteCmd) Run(g globalCmd) error {
	if c.IP == nil && c.Host == nil {
		return errors.New("IP or Host is required")
	}

	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	el, changed := c.Delete(el, c.IP, c.Host)

	if changed {
		return WriteEntries(g.Hosts, el)
	}

	return nil
}

func (c deleteCmd) Delete(el []Entry, ip, host *string) ([]Entry, bool) {
	idx := -1
	for i, e := range el {
		if ip != nil && host != nil && e.IP == *ip && e.Host == *host {
			idx = i
			break
		} else if ip != nil && e.IP == *ip {
			idx = i
			break
		} else if host != nil && e.Host == *host {
			idx = i
			break
		}
	}
	if idx != -1 { // found
		el = append(el[:idx], el[idx+1:]...)
		return el, true
	}
	return el, false
}
