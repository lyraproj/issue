package issue

import "fmt"

func ExampleA_an() {
	fmt.Println(AnOrA(`string`))
	fmt.Println(AnOrA(`integer`))
	// Output:
	// a string
	// an integer
}

func ExampleFirstToLower() {
	fmt.Println(FirstToLower(`MyNameIsBob`))
	fmt.Println(FirstToLower(`_MyNameIsBob`))
	fmt.Println(FirstToLower(`__MyNameIsBob`))
	// Output:
	// myNameIsBob
	// _myNameIsBob
	// __myNameIsBob
}

func ExampleJoinErrors() {
	fmt.Println(JoinErrors([]string{`first`, `second`, `third`}))
	// Output:
	// first
	// second
	// third
}

func ExampleCamelToSnakeCase() {
	fmt.Println(CamelToSnakeCase(`MyNameIsBob`))
	fmt.Println(CamelToSnakeCase(`_MyNameIsBob`))
	fmt.Println(CamelToSnakeCase(`__MyNameIsBob`))
	fmt.Println(CamelToSnakeCase(`SomeID`))
	fmt.Println(CamelToSnakeCase(`SomeIDsomething`))
	fmt.Println(CamelToSnakeCase(`SomeID_OfSomething`))
	// Output:
	// my_name_is_bob
	// _my_name_is_bob
	// __my_name_is_bob
	// some_id
	// some_id_something
	// some_id_of_something
}

func ExampleSnakeToCamelCase() {
	fmt.Println(SnakeToCamelCase(`my_name_is_bob`))
	fmt.Println(SnakeToCamelCase(`_my_name_is_bob`))
	fmt.Println(SnakeToCamelCase(`__my_name_is_bob`))
	// Output:
	// MyNameIsBob
	// _MyNameIsBob
	// __MyNameIsBob
}

func ExampleSnakeToCamelCaseDC() {
	fmt.Println(SnakeToCamelCaseDC(`my_name_is_bob`))
	fmt.Println(SnakeToCamelCaseDC(`_my_name_is_bob`))
	fmt.Println(SnakeToCamelCaseDC(`__my_name_is_bob`))
	// Output:
	// myNameIsBob
	// _myNameIsBob
	// __myNameIsBob
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
