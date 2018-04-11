package issue

import (
	"fmt"
)

// this would be an enum in most other languages
const (
	SEVERITY_IGNORE      = Severity(1)
	SEVERITY_DEPRECATION = Severity(2)
	SEVERITY_WARNING     = Severity(3)
	SEVERITY_ERROR       = Severity(4)
)

// Severity used in reported issues
type Severity int

// String returns the severity in lowercase string form
func (severity Severity) String() string {
	switch severity {
	case SEVERITY_IGNORE:
		return `ignore`
	case SEVERITY_DEPRECATION:
		return `warning`
	case SEVERITY_WARNING:
		return `warning`
	case SEVERITY_ERROR:
		return `error`
	default:
		panic(fmt.Sprintf(`Illegal severity level: %d`, severity))
	}
}

// AssertValid checks that the given severity is one of the recognized severities
func (severity Severity) AssertValid() {
	switch severity {
	case SEVERITY_IGNORE, SEVERITY_DEPRECATION, SEVERITY_WARNING, SEVERITY_ERROR:
		return
	default:
		panic(fmt.Sprintf(`Illegal severity level: %d`, severity))
	}
}
