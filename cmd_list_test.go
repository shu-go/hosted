package main

import (
	"bytes"
	"testing"

	"github.com/shu-go/gotwant"
)

func read(hosts string) []entry {
	input := bytes.NewBufferString(hosts)
	el, _ := readEntriesFromReader(input)
	return el
}

func list(el []entry) string {
	ls := listCmd{}
	buf := bytes.Buffer{}
	ls.list(el, &buf)
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

	t.Run("Comment/Option", func(t *testing.T) {
		el := read(hosts)
		comment := "server"
		ls := listCmd{Comment: &comment}
		buf := bytes.Buffer{}
		ls.list(el, &buf)
		gotwant.Test(t, buf.String(), `# 102.54.94.97 rhino.acme.com # source server
# 192.168.1.201 server01 # new server
`)
	})

	t.Run("Comment/Args", func(t *testing.T) {
		el := read(hosts)
		ls := listCmd{}
		ls.feed([]string{"server"})
		buf := bytes.Buffer{}
		ls.list(el, &buf)
		gotwant.Test(t, buf.String(), `# 102.54.94.97 rhino.acme.com # source server
# 192.168.1.201 server01 # new server
`)
	})

}
