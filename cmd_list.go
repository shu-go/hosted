package main

import (
	"fmt"
)

type listCmd struct {
}

func (c listCmd) Run(g globalCmd) error {
	el, err := ReadEntries(g.Hosts)
	if err != nil {
		return err
	}

	for _, e := range el {
		if e.Type == HostEntry {
			fmt.Println(e)
		}
	}

	return nil
}
