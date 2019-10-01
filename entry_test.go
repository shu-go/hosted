package main

import (
	"testing"

	"github.com/shu-go/gotwant"
)

func TestEntry(t *testing.T) {
	t.Run("Read", func(t *testing.T) {
		cc := []gotwant.TestCase{
			gotwant.Case(ReadEntry(""), &Entry{Type: EmptyEntry}),
			gotwant.Case(ReadEntry("# hoge"), &Entry{Type: CommentEntry, Comment: " hoge"}),
			gotwant.Case(ReadEntry(" # hoge"), &Entry{Type: CommentEntry, Comment: " hoge"}),
			gotwant.Case(ReadEntry("192.168.1.1"), (*Entry)(nil)),
			gotwant.Case(ReadEntry("192.168.1.1 my-host"), &Entry{Type: HostEntry, IP: "192.168.1.1", Host: "my-host", Enabled: true}),
			gotwant.Case(ReadEntry("192.168.1.1 my-host # myhome"), &Entry{Type: HostEntry, IP: "192.168.1.1", Host: "my-host", Comment: "myhome", Enabled: true}),
			gotwant.Case(ReadEntry("# 192.168.1.1 my-host # myhome"), &Entry{Type: HostEntry, IP: "192.168.1.1", Host: "my-host", Comment: "myhome", Enabled: false}),
		}
		gotwant.TestAll(t, cc)
	})
}
