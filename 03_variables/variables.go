// In Go, variables are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	// if not explicitely defined.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// In Go the `:=` syntax is shorthand for declaring
	// and initializing a variable, e.g. for
	// `var f = "apple"` in this case. Go further
	// infers `var f string = "apple"`
	f := "apple"
	fmt.Println(f)
}
