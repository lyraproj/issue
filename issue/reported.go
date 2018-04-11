package issue

import (
	"bytes"
)

// A Reported instance contains information of an issue such as an
// error or a warning. It contains an Issue and arguments needed to
// format that issue. It also contains the location where the issue
// was reported.
type Reported interface {
	// Argument returns the argument for the given key or nil if no
	// such argument exists
	Argument(key string) interface{}

	// Code returns the issue code
	Code() Code

	// Error produces a string from the receives issue and arguments
	Error() string

	// Error produces a string from the receives issue and arguments
	// and writes it to the given buffer
	ErrorTo(*bytes.Buffer)

	// Location returns the location where the issue was reported
	Location() Location

	// OffsetByLocation returns a copy of the receiver where the location
	// is offset by the given location. This is useful when the original
	// source is embedded in a another file.
	OffsetByLocation(location Location) Reported

	// Severity returns the severity
	Severity() Severity

	// String is an alias for Error
	String() string
}

type reported struct {
	issueCode Code
	severity  Severity
	args      H
	location  Location
}

func NewReported(code Code, severity Severity, args H, location Location) Reported {
	return &reported{code, severity, args, location}
}

func (ri *reported) Argument(key string) interface{} {
	return ri.args[key]
}

func (ri *reported) OffsetByLocation(location Location) Reported {
	loc := ri.location
	if loc == nil {
		loc = location
	} else {
		loc = NewLocation(location.File(), location.Line()+loc.Line(), location.Pos())
	}
	return &reported{ri.issueCode, ri.severity, ri.args, loc}
}

func (ri *reported) Error() (str string) {
	b := bytes.NewBufferString(``)
	ri.ErrorTo(b)
	return b.String()
}

func (ri *reported) ErrorTo(b *bytes.Buffer) {
	IssueForCode(ri.issueCode).Format(b, ri.args)
	if ri.location != nil {
		b.WriteByte(' ')
		appendLocation(b, ri.location)
	}
}

func (ri *reported) Location() Location {
	return ri.location
}

func (ri *reported) String() (str string) {
	return ri.Error()
}

func (ri *reported) Code() Code {
	return ri.issueCode
}

func (ri *reported) Severity() Severity {
	return ri.severity
}
