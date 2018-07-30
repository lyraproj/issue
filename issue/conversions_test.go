package issue

import "fmt"

func ExampleCamelToSnakeCase() {
	fmt.Println(CamelToSnakeCase(`MyNameIsBob`))
	// Output: my_name_is_bob
}
