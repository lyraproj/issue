package issue

import (
	"fmt"
	"os"
	"testing"
)

func assertEqual(t *testing.T, e, a interface{}) {
	if e != a {
		t.Errorf("Expected '%s', got '%s'", e, a)
	}
}

func ExampleFprintf() {
	MapFprintf(os.Stdout, "%{foo} %{fee} %{fum}", H{"foo": "hello", "fee": "great", "fum": "world"})
	// Output: hello great world
}

func ExampleFprintf_ignoredFlags() {
	MapFprintf(os.Stdout, "%{foo}4d, %{foo}o, %{foo}X", H{"foo": 23})
	// Output: 234d, 23o, 23X
}

func ExampleFprintf_flags() {
	MapFprintf(os.Stdout, "%<foo>4d, %<foo>o, %<foo>X", H{"foo": 23})
	// Output:   23, 27, 17
}

func ExampleMapFprintf_duplicateArguments() {
	MapFprintf(os.Stdout, "%{foo} %{fee} %{foo}", H{"foo": "boys", "fee": "will be"})
	// Output: boys will be boys
}

func ExampleMapFprintf_missingKey() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	MapFprintf(os.Stdout, "%{foo} %{fee} %{fum}", H{"foo": "hello", "fum": "world"})
	// Output: hello %!{fee}(MISSING) world
}
