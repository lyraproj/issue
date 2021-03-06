package issue

import "fmt"

func ExampleNewReported() {
	// Issues are normally declared in an init() function and they end up in a global
	// variable. The withIssues() function used here is only for test purposes
	withIssues(func() {
		const (
			FirstIssue  = `FIRST_ISSUE`
			SecondIssue = `SECOND_ISSUE`
		)

		// Issue using %{name} notation to represent a value with default format (%v) and
		// %<> notation to use a specific format (here %T)
		Hard(FirstIssue, "The %{item} is of incorrect type. Expected int32, got %<value>T")

		// Issue using an ArgumentsFormatter to prefix the item with a capitalized article
		Hard2(SecondIssue, "%{item} cannot be used here", HF{`item`: UcAnOrA})

		err1 := NewReported(FirstIssue, SeverityError, H{`item`: `width`, `value`: int16(12)}, NewLocation(`/tmp/test`, 32, 14))

		err2 := NewReported(SecondIssue, SeverityError, H{`item`: `integer`}, NewLocation(`/tmp/test`, 42, 8))

		fmt.Println(err1)
		fmt.Println(err2)

	})
	// Output:
	// The width is of incorrect type. Expected int32, got int16 (file: /tmp/test, line: 32, column: 14)
	// An integer cannot be used here (file: /tmp/test, line: 42, column: 8)
}
