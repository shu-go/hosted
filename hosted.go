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
		/*
			switch e.Type {
			case EmptyEntry:
				content = append(content, "\n"...)

			case CommentEntry:
				content = append(content, "# "...)
				content = append(content, e.Comment...)
				content = append(content, "\n"...)

			case HostEntry:
				if !e.Enabled {
					content = append(content, "# "...)
				}
				content = append(content, e.IP...)
				content = append(content, "  "...)
				content = append(content, e.Host...)
				if e.Comment != "" {
					content = append(content, "  # "...)
					content = append(content, e.Comment...)
				}
				content = append(content, "\n"...)
			}
		*/
	}

	return ioutil.WriteFile(filename, content, os.ModePerm)
}

func main() {
	app := gli.NewWith(&globalCmd{})
	app.Name = "hosted"
	app.Desc = ""
	app.Version = "0.1.0"
	app.Usage = `edit hosts file`
	app.Copyright = "(C) 2019 Shuhei Kubota"
	app.Run(os.Args)
}
