package main

import (
	"errors"
)

type disableCmd struct {
	_ struct{} `help:"change normal entry -> comment if exists"`

	IP      *string
	Host    *string
	Comment *string
}

func (c disableCmd) Run(g globalCmd) error {
	if c.IP == nil && c.Host == nil && c.Comment == nil {
		return errors.New("IP or Host or Comment is required")
	}

	el, err := readEntries(g.Hosts)
	if err != nil {
		return err
	}

	el = c.disable(el)

	if el != nil {
		return writeEntries(g.Hosts, el)
	}
	return nil
}

func (c disableCmd) disable(el []entry) []entry {
	dirty := false

	for i, e := range el {
		found := false
		if matches(e, c.IP, c.Host, c.Comment) {
			found = true
		}

		if found && e.Type == hostEntry && el[i].Enabled {
			el[i].Enabled = false
			dirty = true
		}
	}

	if dirty {
		return el
	}
	return nil
}
