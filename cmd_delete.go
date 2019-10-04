package main

import (
	"errors"

	"github.com/shu-go/clise"
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

	el, changed := c.DeleteFrom(el)

	if changed {
		return WriteEntries(g.Hosts, el)
	}

	return nil
}

func (c deleteCmd) DeleteFrom(el []Entry) ([]Entry, bool) {
	changed := false
	clise.Filter(&el, func(i int) bool {
		if c.IP != nil && c.Host != nil && el[i].IP == *c.IP && el[i].Host == *c.Host {
			changed = true
			return false
		} else if c.IP != nil && el[i].IP == *c.IP {
			changed = true
			return false
		} else if c.Host != nil && el[i].Host == *c.Host {
			changed = true
			return false
		}
		return true
	})
	return el, changed
}
