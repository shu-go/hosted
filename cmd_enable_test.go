package main

import (
	"testing"

	"github.com/shu-go/gotwant"
)

func TestEnable(t *testing.T) {
	hosts := `
# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
# localhost name resolution is handle within DNS itself.

127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`

	t.Run("IP", func(t *testing.T) {
		ip := "192.168.1.201"
		el := read(hosts)
		enable := enableCmd{IP: &ip}
		el = enable.enable(el)
		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
192.168.1.201 server01 # new server
`)

		// no matches

		ip = "111.111.1.111"
		el = read(hosts)
		enable = enableCmd{IP: &ip}
		el = enable.enable(el)
		result = list(el)
		gotwant.Test(t, result, ``)
	})

	t.Run("Host", func(t *testing.T) {
		host := "server01"
		el := read(hosts)
		enable := enableCmd{Host: &host}
		el = enable.enable(el)
		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
192.168.1.201 server01 # new server
`)

		// matches, but not changed

		host = "localhost"
		el = read(hosts)
		enable = enableCmd{Host: &host}
		el = enable.enable(el)
		result = list(el)
		gotwant.Test(t, result, ``)
	})

}
