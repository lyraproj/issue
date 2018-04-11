package issue

import (
	"fmt"
	"bytes"
)

var NO_ARGS = H{}

type Code string

type ArgFormatter func(value interface{}) string

type HF map[string]ArgFormatter

type Issue interface {
	// ArgFormatters returns the argument formatters or nil if no such
	// formatters exists
	ArgFormatters() HF

	// Code returns the issue code
	Code() Code

	// Format uses the receivers format string and the given arguments to
	// write a string onto the given buffer
	Format(b *bytes.Buffer, arguments H)

	// IsDemotable returns false for soft issues and true for hard issues
	IsDemotable() bool

	// MessageFormat returns the format used when formatting the receiver
	MessageFormat() string
}

type issue struct {
	code          Code
	messageFormat string
	argFormats    HF
	demotable     bool
}

var issues = map[Code]*issue{}

func Hard(code Code, messageFormat string) Issue {
	return addIssue(code, messageFormat, false, nil)
}

func Hard2(code Code, messageFormat string, argFormats HF) Issue {
	return addIssue(code, messageFormat, false, argFormats)
}

func SoftIssue(code Code, messageFormat string) Issue {
	return addIssue(code, messageFormat, true, nil)
}

func SoftIssue2(code Code, messageFormat string, argFormats HF) Issue {
	return addIssue(code, messageFormat, true, argFormats)
}

func (issue *issue) ArgFormatters() HF {
	return issue.argFormats
}

func (issue *issue) Code() Code {
	return issue.code
}

func (issue *issue) IsDemotable() bool {
	return issue.demotable
}

func (issue *issue) MessageFormat() string {
	return issue.messageFormat
}

func (issue *issue) Format(b *bytes.Buffer, arguments H) {
	var args H
	af := issue.ArgFormatters()
	if af != nil {
		args = make(H, len(arguments))
		for k, v := range arguments {
			if a, ok := af[k]; ok {
				v = a(v)
			}
			args[k] = v
		}
	} else {
		args = arguments
	}
	MapFprintf(b, issue.MessageFormat(), args)
}

func addIssue(code Code, messageFormat string, demotable bool, argFormats HF) Issue {
	dsc := &issue{code, messageFormat, argFormats, demotable}
	issues[code] = dsc
	return dsc
}

// Returns the Issue for a Code. Will panic if the given code does not represent
// an existing issue
func IssueForCode(code Code) Issue {
	if dsc, ok := issues[code]; ok {
		return dsc
	}
	panic(fmt.Sprintf("internal error: no issue found for issue code '%s'", code))
}

func IssueForCode2(code Code) (dsc Issue, ok bool) {
	dsc, ok = issues[code]
	return
}
