package main

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/shu-go/gli"
)

var Version string

func init() {
	if Version == "" {
		Version = "dev-" + time.Now().Format("20060102")
	}
}

type globalCmd struct {
	Hosts string `cli:"hosts=PATH"  default:"C:\\Windows\\System32\\Drivers\\etc\\hosts"`

	List   listCmd `cli:"list,ls"`
	Add    addCmd
	Delete deleteCmd `cli:"del,delete"`

	Enable  enableCmd  `cli:"enable,o"`
	Disable disableCmd `cli:"disable,x"`
}

func ReadEntries(filename string) ([]Entry, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadEntriesFromReader(f)
}

func ReadEntriesFromReader(in io.Reader) ([]Entry, error) {
	el := make([]Entry, 0, 10)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		e := ReadEntry(scanner.Text())
		if e != nil {
			el = append(el, *e)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return el, nil
}

func WriteEntries(filename string, el []Entry) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return WriteEntriesToWriter(f, el)
}

func WriteEntriesToWriter(out io.Writer, el []Entry) error {
	content := make([]byte, 0, 1024)

	for _, e := range el {
		content = append(content, e.String()...)
		content = append(content, "\n"...)
	}

	_, err := io.WriteString(out, string(content))
	return err
}

func matches(e Entry, ip, host *string) bool {
	if ip != nil && e.IP != *ip {
		return false
	}
	if host != nil && e.Host != *host {
		return false
	}

	return true
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
	app.Run(os.Args)
}
