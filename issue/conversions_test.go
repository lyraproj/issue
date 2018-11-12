package issue

import "fmt"

func ExampleCamelToSnakeCase() {
	fmt.Println(CamelToSnakeCase(`MyNameIsBob`))
	// Output: my_name_is_bob
}

func ExampleSnakeToCamelCase() {
	fmt.Println(SnakeToCamelCase(`my_name_is_bob`))
	// Output: MyNameIsBob
}

func ExampleUnindent() {
	fmt.Println(Unindent(`
     No whitespace in front of this line.
       Two whitespaces in front of this one
     `))

	// Output:
	// No whitespace in front of this line.
	//   Two whitespaces in front of this one
}
