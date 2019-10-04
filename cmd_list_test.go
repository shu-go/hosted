package main

import (
	"bytes"
	"testing"

	"github.com/shu-go/gotwant"
)

func read(hosts string) []Entry {
	input := bytes.NewBufferString(hosts)
	el, _ := ReadEntriesFromReader(input)
	return el
}

func list(el []Entry) string {
	ls := listCmd{}
	buf := bytes.Buffer{}
	ls.List(el, &buf)
	return buf.String()
}

func TestCmdList(t *testing.T) {
	hosts := `
# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
# localhost name resolution is handle within DNS itself.

127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`

	result := list(read(hosts))
	gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)

}
