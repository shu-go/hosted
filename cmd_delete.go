package main

import (
	"errors"
	"strings"

	"github.com/shu-go/clise"
)

type deleteCmd struct {
	_ struct{} `help:"delete an entry with  --ip 192.168.1.200 and/or --host oldserver"`

	IP      *string
	Host    *string
	Comment *string
}

func (c deleteCmd) Run(g globalCmd) error {
	if c.IP == nil && c.Host == nil && c.Comment == nil {
		return errors.New("IP or Host or Comment is required")
	}

	el, err := readEntries(g.Hosts)
	if err != nil {
		return err
	}

	el, changed := c.deleteFrom(el)

	if changed {
		return writeEntries(g.Hosts, el)
	}

	return nil
}

func (c deleteCmd) deleteFrom(el []entry) ([]entry, bool) {
	changed := false

	clise.Filter(&el, func(i int) bool {
		e := el[i]

		deleting := true
		if c.IP != nil {
			deleting = deleting && (e.IP == *c.IP)
		}
		if c.Host != nil {
			deleting = deleting && (e.Host == *c.Host)
		}
		if c.Comment != nil {
			deleting = deleting && strings.Contains(e.Comment, *c.Comment)
		}

		if deleting {
			changed = true
			return false
		}
		return true
	})
	return el, changed
}
