package main

import (
	"fmt"
	"io"
	"os"
)

type listCmd struct {
}

func (c listCmd) Run(g globalCmd) error {
	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	c.List(el, os.Stdout)

	return nil
}

func (c listCmd) List(el []Entry, out io.Writer) {
	for _, e := range el {
		if e.Type == HostEntry {
			fmt.Fprintln(out, e)
		}
	}
}
