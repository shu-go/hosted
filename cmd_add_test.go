package main

import (
	"testing"

	"github.com/shu-go/gotwant"
)

func TestCmdAdd(t *testing.T) {
	hosts := `
# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
# localhost name resolution is handle within DNS itself.

127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`

	t.Run("Option", func(t *testing.T) {
		el := read(hosts)
		add := addCmd{
			IP:      "::2",
			Host:    "hogeserver",
			Comment: "this is a server",
		}
		el = add.AddTo(el)

		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
::2 hogeserver # this is a server
`)
	})

	t.Run("Args", func(t *testing.T) {
		el := read(hosts)
		add := addCmd{}
		add.Feed([]string{"::2", "hogeserver", "this is a server"})
		el = add.AddTo(el)

		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
::2 hogeserver # this is a server
`)
	})

}
