package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/shu-go/gli/v2"
)

// Version is app version
var Version string

type globalCmd struct {
	Hosts string `cli:"hosts=PATH"  default:"C:\\Windows\\System32\\Drivers\\etc\\hosts"`

	List   listCmd `cli:"list,ls"`
	Add    addCmd
	Delete deleteCmd `cli:"del,delete"`

	Enable  enableCmd  `cli:"enable,o"`
	Disable disableCmd `cli:"disable,x"`
}

func readEntries(filename string) ([]entry, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readEntriesFromReader(f)
}

func readEntriesFromReader(in io.Reader) ([]entry, error) {
	el := make([]entry, 0, 10)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		e := readEntry(scanner.Text())
		if e != nil {
			el = append(el, *e)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return el, nil
}

func writeEntries(filename string, el []entry) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return writeEntriesToWriter(f, el)
}

func writeEntriesToWriter(out io.Writer, el []entry) error {
	content := make([]byte, 0, 1024)

	for _, e := range el {
		content = append(content, e.String()...)
		content = append(content, "\r\n"...)
	}

	_, err := io.WriteString(out, string(content))
	return err
}

func matches(e entry, ip, host, comment *string) bool {
	if ip == nil && host == nil && comment == nil {
		return false
	}

	matching := true

	if ip != nil {
		matching = matching && (e.IP == *ip)
	}
	if host != nil {
		matching = matching && (e.Host == *host)
	}
	if comment != nil {
		matching = matching && strings.Contains(e.Comment, *comment)
	}

	return matching
}

func main() {
	app := gli.NewWith(&globalCmd{})
	app.Name = "hosted"
	app.Desc = "edit Windows HOSTS file"
	app.Version = Version
	app.Usage = `RUN AS ADMINISTRATOR

# ADD new server "server01" as 192.168.1.201
hosted add --ip 192.168.1.201 --host server01 --comment "new server"
hosted add 192.168.1.201 server01 new server

# REMOVE a wrong entry
hosted delete --ip 192.168.1.210

# COMMENT-OUT (disable) oldserver
hosted disable --host oldserver
hosted x --host oldserver

# COMMENT-IN (enable) oldserver back
hosted enable --host oldserver
hosted o --host oldserver
`
	app.Copyright = "(C) 2019 Shuhei Kubota"
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
