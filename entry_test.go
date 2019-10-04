package main

import (
	"testing"

	"github.com/shu-go/gotwant"
)

func TestEntry(t *testing.T) {
	t.Run("Read", func(t *testing.T) {
		cc := []gotwant.TestCase{
			gotwant.Case(readEntry(""), &entry{Type: emptyEntry}),
			gotwant.Case(readEntry("# hoge"), &entry{Type: commentEntry, Comment: " hoge"}),
			gotwant.Case(readEntry(" # hoge"), &entry{Type: commentEntry, Comment: " hoge"}),
			gotwant.Case(readEntry("192.168.1.1"), (*entry)(nil)),
			gotwant.Case(readEntry("192.168.1.1 my-host"), &entry{Type: hostEntry, IP: "192.168.1.1", Host: "my-host", Enabled: true}),
			gotwant.Case(readEntry("192.168.1.1 my-host # myhome"), &entry{Type: hostEntry, IP: "192.168.1.1", Host: "my-host", Comment: "myhome", Enabled: true}),
			gotwant.Case(readEntry("# 192.168.1.1 my-host # myhome"), &entry{Type: hostEntry, IP: "192.168.1.1", Host: "my-host", Comment: "myhome", Enabled: false}),
		}
		gotwant.TestAll(t, cc)
	})
}
