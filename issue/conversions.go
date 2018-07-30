package issue

import (
	"bytes"
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
