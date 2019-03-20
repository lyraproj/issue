package issue

import (
	"bytes"
	"strconv"
)

type Location interface {
	File() string

	Line() int

	// Position on line
	Pos() int
}

type Located interface {
	Location() Location
}

type location struct {
	file string
	line int
	pos  int
}

func NewLocation(file string, line, pos int) Location {
	return &location{file, line, pos}
}

func (l *location) File() string {
	return l.file
}

func (l *location) Line() int {
	return l.line
}

func (l *location) Pos() int {
	return l.pos
}

func LocationString(location Location) string {
	b := bytes.NewBufferString(``)
	appendLocation(b, location)
	return b.String()
}

func appendLocation(b *bytes.Buffer, location Location) {
	if location == nil {
		return
	}
	file := location.File()
	line := location.Line()
	if file == `` && line <= 0 {
		return
	}

	pos := location.Pos()
	b.WriteByte('(')
	if file != `` {
		b.WriteString(`file: `)
		b.WriteString(file)
		if line > 0 {
			b.WriteString(`, `)
		}
	}
	if line > 0 {
		b.WriteString(`line: `)
		b.WriteString(strconv.Itoa(line))
		if pos > 0 {
			b.WriteString(`, column: `)
			b.WriteString(strconv.Itoa(pos))
		}
	}
	b.WriteByte(')')
}
