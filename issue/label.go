package issue

import (
	"fmt"
)

// A Labeled is an object that has a label in the form of a string
type Labeled interface {
	// Returns a very brief description of this expression suitable to use in error messages
	Label() string
}

// A Named is an object that has a name in the form of a string
type Named interface {
	Name() string
}

// Label returns the Label for a Labeled argument, the Name for a Named argument, a string
// verbatim, or the resulting text from doing a Sprintf with "value of type %T" for other
// types of arguments.
func Label(e interface{}) string {
	if l, ok := e.(Labeled); ok {
		return l.Label()
	}
	if n, ok := e.(Named); ok {
		return n.Name()
	}
	if s, ok := e.(string); ok {
		return s
	}
	return fmt.Sprintf(`value of type %T`, e)
}

// A_an returns the non capitalized article for the label of the given argument
func A_an(e interface{}) string {
	label := Label(e)
	return fmt.Sprintf(`%s %s`, Article(label), label)
}

// A_anUc returns the capitalized article for the label of the given argument
func A_anUc(e interface{}) string {
	label := Label(e)
	return fmt.Sprintf(`%s %s`, ArticleUc(label), label)
}

// Article returns the non capitalized article for the given string
func Article(s string) string {
	if s == `` {
		return `a`
	}
	switch s[0] {
	case 'A', 'E', 'I', 'O', 'U', 'Y', 'a', 'e', 'i', 'o', 'u', 'y':
		return `an`
	default:
		return `a`
	}
}

// ArticleUc returns the capitalized article for the given string
func ArticleUc(s string) string {
	if s == `` {
		return `A`
	}
	switch s[0] {
	case 'A', 'E', 'I', 'O', 'U', 'Y', 'a', 'e', 'i', 'o', 'u', 'y':
		return `An`
	default:
		return `A`
	}
}
