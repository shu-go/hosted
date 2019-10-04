package main

import (
	"testing"

	"github.com/shu-go/gotwant"
)

func TestCmdDelete(t *testing.T) {
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
		el := read(hosts)
		ip := "192.168.1.201"
		del := deleteCmd{
			IP: &ip,
		}
		el, changed := del.deleteFrom(el)

		gotwant.Test(t, changed, true)

		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
`)

		el = read(hosts)
		ip = "::1"
		del = deleteCmd{
			IP: &ip,
		}
		el, changed = del.deleteFrom(el)

		gotwant.Test(t, changed, true)

		result = list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)
	})

	t.Run("Host", func(t *testing.T) {
		el := read(hosts)
		host := "localhost"
		del := deleteCmd{
			Host: &host,
		}
		el, changed := del.deleteFrom(el)

		gotwant.Test(t, changed, true)

		result := list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)

		el = read(hosts)
		host = "hgoehgoelocalhost"
		del = deleteCmd{
			Host: &host,
		}
		el, changed = del.deleteFrom(el)

		gotwant.Test(t, changed, false)

		result = list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)
	})

	t.Run("Comment", func(t *testing.T) {
		el := read(hosts)
		comment := "server"
		del := deleteCmd{
			Comment: &comment,
		}
		el, changed := del.deleteFrom(el)

		gotwant.Test(t, changed, true)

		result := list(el)
		gotwant.Test(t, result, `# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
`)

		el = read(hosts)
		comment = "hgoehgoelocalhost"
		del = deleteCmd{
			Comment: &comment,
		}
		el, changed = del.deleteFrom(el)

		gotwant.Test(t, changed, false)

		result = list(el)
		gotwant.Test(t, result, `# 102.54.94.97 rhino.acme.com # source server
# 38.25.63.10 x.acme.com # x client host
127.0.0.1 localhost # THIS IS LOCALHOST
::1 localhost
# 118.151.235.191 example.com
# 192.168.1.201 server01 # new server
`)
	})
}
