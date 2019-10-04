package main

import (
	"fmt"
	"io"
	"os"
)

type listCmd struct {
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
		if e.Type == hostEntry {
			fmt.Fprintln(out, e)
		}
	}
}
