package main

import "regexp"

type EntryType int

const (
	EmptyEntry EntryType = iota
	HostEntry
	CommentEntry
)

func (t EntryType) String() string {
	switch t {
	case EmptyEntry:
		return "_"
	case HostEntry:
		return "H"
	case CommentEntry:
		return "#"
	}
	return ""
}

type Entry struct {
	Type EntryType

	IP      string
	Host    string
	Comment string
	Enabled bool
}

func (e Entry) String() string {
	switch e.Type {
	case CommentEntry:
		return "#" + e.Comment
	case HostEntry:
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
var hostRE = regexp.MustCompile(`^(#?)\s*([0-9\.:]{1,})\s*(\S+)\s*(?:#\s*(.+))?`)
var commentRE = regexp.MustCompile(`^#(.*)`)

func ReadEntry(line string) *Entry {
	if line == "" {
		return &Entry{
			Type: EmptyEntry,
		}
	}

	re := hostRE.Copy()
	subs := re.FindStringSubmatch(line)
	if len(subs) != 0 {
		return &Entry{
			Type:    HostEntry,
			IP:      subs[2],
			Host:    subs[3],
			Comment: subs[4],
			Enabled: (subs[1] != "#"),
		}
	}

	re = commentRE.Copy()
	subs = re.FindStringSubmatch(line)
	if len(subs) != 0 {
		return &Entry{
			Type:    CommentEntry,
			Comment: subs[1],
		}
	}

	return nil
}
