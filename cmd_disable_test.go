package main

import (
	"testing"

	"github.com/shu-go/gotwant"
)

func TestDisable(t *testing.T) {
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
		ip := "127.0.0.1"
		el := read(hosts)
		disable := disableCmd{IP: &ip}
		el = disable.disable(el)
		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
# 127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)

		// no matches

		ip = "111.111.1.111"
		el = read(hosts)
		disable = disableCmd{IP: &ip}
		el = disable.disable(el)
		result = list(el)
		gotwant.Test(t, result, ``)
	})

	t.Run("Host", func(t *testing.T) {
		host := "localhost"
		el := read(hosts)
		disable := disableCmd{Host: &host}
		el = disable.disable(el)
		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
# 127.0.0.1 localhost # THIS IS LOCALHOST
# ::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)

		// matches, but not changed

		host = "server01"
		el = read(hosts)
		disable = disableCmd{Host: &host}
		el = disable.disable(el)
		result = list(el)
		gotwant.Test(t, result, ``)
	})

	t.Run("Comment", func(t *testing.T) {
		comment := "HOST"
		el := read(hosts)
		disable := disableCmd{Comment: &comment}
		el = disable.disable(el)
		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
# 127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)
	})

}
