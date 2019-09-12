package main

import (
	"errors"
	"fmt"
	"strings"
)

type addCmd struct {
	_                 struct{} `help:"{IP} {Host} [Comment] "`
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

	found := false
	for i := range el {
		e := el[i]

		fmt.Println(e)
		if e.Type == HostEntry && (e.IP == c.IP || e.Host == c.Host) {
			fmt.Println("  FOUND")
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

	return WriteEntries(g.Hosts, el)
}
