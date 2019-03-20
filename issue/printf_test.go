package issue

import (
	"fmt"
	"os"
)

func ExampleMapFprintf() {
	_, _ = MapFprintf(os.Stdout, "%{foo} %{fee} %{fum}", H{"foo": "hello", "fee": "great", "fum": "world"})
	// Output: hello great world
}

func ExampleMapSprintf() {
	fmt.Print(MapSprintf("%{foo} %{fee} %{fum}\n", H{"foo": "hello", "fee": "great", "fum": "world"}))
	// Output: hello great world
}

func ExampleMapFprintf_ignoredFlags() {
	_, _ = MapFprintf(os.Stdout, "%{foo}4d, %{foo}o, %{foo}X", H{"foo": 23})
	// Output: 234d, 23o, 23X
}

func ExampleMapFprintf_flags() {
	_, _ = MapFprintf(os.Stdout, "%<foo>4d, %<foo>o, %<foo>X", H{"foo": 23})
	// Output:   23, 27, 17
}

func ExampleMapFprintf_duplicateArguments() {
	_, _ = MapFprintf(os.Stdout, "%{foo} %{fee} %{foo}", H{"foo": "boys", "fee": "will be"})
	// Output: boys will be boys
}

func ExampleMapFprintf_missingKey() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	_, _ = MapFprintf(os.Stdout, "%{foo} %{fee} %{fum}", H{"foo": "hello", "fum": "world"})
	// Output: hello %!{fee}(MISSING) world
}
