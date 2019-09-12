package main

import "errors"

type deleteCmd struct {
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

	idx := -1
	for i, e := range el {
		if c.IP != nil && c.Host != nil && e.IP == *c.IP && e.Host == *c.Host {
			idx = i
			break
		} else if c.IP != nil && e.IP == *c.IP {
			idx = i
			break
		} else if c.Host != nil && e.Host == *c.Host {
			idx = i
			break
		}
	}
	if idx != -1 { // found
		el = append(el[:idx], el[idx+1:]...)
		return WriteEntries(g.Hosts, el)
	}
	return nil
}
