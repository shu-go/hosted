package main

import (
	"errors"
	"strings"
)

type addCmd struct {
	_                 struct{} `help:"add an entry with args: {IP} {Host} [Comment] "`
	IP, Host, Comment string
}

func (c addCmd) Run(args []string, g globalCmd) error {
	err := c.Feed(args)
	if err != nil {
		return err
	}

	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	el = c.AddTo(el)

	return WriteEntries(g.Hosts, el)
}

func (c *addCmd) Feed(args []string) error {
	idx := 0
	if c.IP == "" {
		if len(args) < idx+1 {
			return errors.New("no IP")
		}
		c.IP = args[idx]
		idx++
	}
	if c.Host == "" {
		if len(args) < idx+1 {
			return errors.New("no Host")
		}
		c.Host = args[idx]
		idx++
	}
	if c.Comment == "" {
		if len(args) < idx+1 {
			//noerror
			//return errors.New("no Comment")
		} else {
			c.Comment = strings.Join(args[idx:], " ")
		}
	}
	return nil
}

func (c addCmd) AddTo(el []Entry) []Entry {
	found := false
	for i := range el {
		e := el[i]

		if e.Type == HostEntry && (e.IP == c.IP || e.Host == c.Host) {
			el[i].IP = c.IP
			el[i].Host = c.Host
			el[i].Comment = c.Comment
			el[i].Enabled = true

			found = true
			break
		}
	}

	if !found {
		el = append(el, Entry{
			Type:    HostEntry,
			IP:      c.IP,
			Host:    c.Host,
			Comment: c.Comment,
			Enabled: true,
		})
	}

	return el
}
