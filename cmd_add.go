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

	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	el = c.Add(el, c.IP, c.Host, c.Comment)

	return WriteEntries(g.Hosts, el)
}

func (c addCmd) Add(el []Entry, ip, host, comment string) []Entry {
	found := false
	for i := range el {
		e := el[i]

		if e.Type == HostEntry && (e.IP == ip || e.Host == host) {
			el[i].IP = ip
			el[i].Host = host
			el[i].Comment = comment
			el[i].Enabled = true

			found = true
			break
		}
	}

	if !found {
		el = append(el, Entry{
			Type:    HostEntry,
			IP:      ip,
			Host:    host,
			Comment: comment,
			Enabled: true,
		})
	}

	return el
}
