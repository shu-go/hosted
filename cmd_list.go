package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type listCmd struct {
	Comment *string
}

func (c listCmd) Run(g globalCmd) error {
	el, err := readEntries(g.Hosts)
	if err != nil {
		return err
	}

	c.list(el, os.Stdout)

	return nil
}

func (c listCmd) list(el []entry, out io.Writer) {
	for _, e := range el {
		if e.Type == hostEntry &&
			(c.Comment == nil || strings.Contains(e.Comment, *c.Comment)) {
			//
			fmt.Fprintln(out, e)
		}
	}
}

func (c *listCmd) feed(args []string) {
	if len(args) > 0 {
		a := args[0]
		c.Comment = &a
	}
}
