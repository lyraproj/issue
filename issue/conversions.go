package issue

import (
	"bytes"
	"math"
	"unicode"
	"fmt"
)

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

// CamelToSnakeCase converts a camel cased name like "NameIsBob" to
// its corresponding snake cased "name_is_bob"
func CamelToSnakeCase(name string) string {
	b := bytes.NewBufferString(``)
	for i, c := range name {
		if unicode.IsUpper(c) {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(c))
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// SnakeToCamelCase converts a snake cased name like "name_is_bob" to
// its corresponding camel cased "NameIsBob"
func SnakeToCamelCase(name string) string {
	b := bytes.NewBufferString(``)

	nextUpper := true
	for _, c := range name {
		if c == '_' {
			nextUpper = true
			continue
		}
		if nextUpper {
			c = unicode.ToUpper(c)
			nextUpper = false
		}
		b.WriteRune(c)
	}
	return b.String()
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

func JoinErrors(e interface{}) string {
	b := bytes.NewBufferString(``)
	for _, error := range e.([]Reported) {
		b.WriteString("\n")
		error.ErrorTo(b)
	}
	return b.String()
}

// Unindent determines the maximum indent that can be stripped by looking at leading whitespace on all lines. Lines that
// consists entirely of whitespace are not included in the computation.
// Strips first line if it's empty, then strips the computed indent from all lines and returns the result.
//
func Unindent(str string) string {
	minIndent := computeIndent(str)
	if minIndent == 0 {
		return str
	}
	r := bytes.NewBufferString(str)
	b := bytes.NewBufferString("")
	first := true
	for {
		line, err := r.ReadString('\n')
		if first {
			first = false
			if line == "\n" {
				continue
			}
		}
		if len(line) > minIndent {
			b.WriteString(line[minIndent:])
		} else if err == nil {
			b.WriteByte('\n')
		} else {
			break
		}
	}
	return b.String()
}

func computeIndent(str string) int {
	minIndent := math.MaxInt64
	r := bytes.NewBufferString(str)
	for minIndent > 0 {
		line, err := r.ReadString('\n')
		ll := len(line)

		for wsCount := 0; wsCount < minIndent && wsCount < ll; wsCount++ {
			c := line[wsCount]
			if c != ' ' && c != '\t' {
				if c != '\n' {
					minIndent = wsCount
				}
				break
			}
		}
		if err != nil {
			break
		}
	}
	if minIndent == math.MaxInt64 {
		minIndent = 0
	}
	return minIndent
}
