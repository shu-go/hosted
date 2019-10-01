package main

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/shu-go/gli"
)

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

	el := make([]Entry, 0, 10)

	scanner := bufio.NewScanner(f)
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
	content := make([]byte, 0, 1024)

	for _, e := range el {
		content = append(content, e.String()...)
		content = append(content, "\n"...)
	}

	return ioutil.WriteFile(filename, content, os.ModePerm)
}

func main() {
	app := gli.NewWith(&globalCmd{})
	app.Name = "hosted"
	app.Desc = "edit Windows HOSTS file"
	app.Version = "0.1.1"
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
