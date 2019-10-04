package main

import "regexp"

type entryType int

const (
	emptyEntry entryType = iota
	hostEntry
	commentEntry
)

func (t entryType) String() string {
	switch t {
	case emptyEntry:
		return "_"
	case hostEntry:
		return "H"
	case commentEntry:
		return "#"
	}
	return ""
}

type entry struct {
	Type entryType

	IP      string
	Host    string
	Comment string
	Enabled bool
}

func (e entry) String() string {
	switch e.Type {
	case commentEntry:
		return "#" + e.Comment
	case hostEntry:
		var s string
		if !e.Enabled {
			s = "# "
		}
		s += e.IP + " " + e.Host
		if e.Comment != "" {
			s += " # " + e.Comment
		}
		return s
	}
	return ""
}

//var hostRE = regexp.MustCompile(`^(#?)\s*(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s*(\S+)\s*(?:#\s*(.+))?`)
var hostRE = regexp.MustCompile(`^\s*(#?)\s*([0-9\.:]{1,})\s+(\S+)\s*(?:#\s*(.+))?`)
var commentRE = regexp.MustCompile(`^\s*#(.*)`)

func readEntry(line string) *entry {
	if line == "" {
		return &entry{
			Type: emptyEntry,
		}
	}

	re := hostRE
	subs := re.FindStringSubmatch(line)
	if len(subs) != 0 {
		return &entry{
			Type:    hostEntry,
			IP:      subs[2],
			Host:    subs[3],
			Comment: subs[4],
			Enabled: (subs[1] != "#"),
		}
	}

	re = commentRE
	subs = re.FindStringSubmatch(line)
	if len(subs) != 0 {
		return &entry{
			Type:    commentEntry,
			Comment: subs[1],
		}
	}

	return nil
}
