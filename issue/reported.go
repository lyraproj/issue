package issue

import (
	"bytes"
	"runtime"
	"strconv"
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
	stack     []runtime.Frame
}

var includeStacktrace = false

// IncludeStacktrace can be set to true to get all Reported to include a stacktrace.
func IncludeStacktrace(flag bool) {
	includeStacktrace = flag
}

func NewReported(code Code, severity Severity, args H, locOrSkip interface{}) Reported {
	var location Location
	skip := 2 // Default is to skip runtime.Callers and this function only
	switch locOrSkip := locOrSkip.(type) {
	case int:
		skip = locOrSkip
	case Location:
		location = locOrSkip
	}

	r := &reported{code, severity, args, location, nil}
	if includeStacktrace {
		// Ask runtime.Callers for up to 100 pcs, including runtime.Callers itself.
		pc := make([]uintptr, 100)
		n := runtime.Callers(skip, pc)
		if n > 0 {
			pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
			frames := runtime.CallersFrames(pc)
			stack := make([]runtime.Frame, 0, n)

			// Loop to get frames.
			// A fixed number of pcs can expand to an indefinite number of Frames.
			for {
				if frame, more := frames.Next(); more {
					stack = append(stack, frame)
				} else {
					break
				}
			}
			r.stack = stack
		}
	}

	if r.location == nil {
		if r.stack == nil {
			for {
				// Use first caller we can find with regards to given skip and use it
				// as the location
				skip--
				if _, f, l, ok := runtime.Caller(skip); ok {
					r.location = NewLocation(f, l, 0)
					break
				}
			}
		} else {
			// Set location to first stack entry
			tf := r.stack[0]
			r.location = NewLocation(tf.File, tf.Line, 0)
		}
	}
	return r
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
	return &reported{ri.issueCode, ri.severity, ri.args, loc, ri.stack}
}

func (ri *reported) Error() (str string) {
	b := bytes.NewBufferString(``)
	ri.ErrorTo(b)
	return b.String()
}

func (ri *reported) ErrorTo(b *bytes.Buffer) {
	IssueForCode(ri.issueCode).Format(b, ri.args)
	if ri.stack != nil {
		for _, f := range ri.stack {
			b.WriteString("\n at ")
			b.WriteString(f.File)
			b.WriteByte(':')
			b.WriteString(strconv.Itoa(f.Line))
			if f.Function != `` {
				b.WriteString(" (")
				b.WriteString(f.Function)
				b.WriteByte(')')
			}
		}
	} else if ri.location != nil {
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
