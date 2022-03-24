// This work was originally from Mark McGranaghan at (https://github.com/mmcgrana/gobyexample).
// And licensed under a Creative Commons Attribution 3.0 Unported License (http://creativecommons.org/licenses/by/3.0/)
// It has been used to provide an example base for multiple languages to provide a basis of comparitive programming
// language study on sytax and simplicity of the languages as far as number of lines of code and operations required
// to perform the same task, as well as compile and run speed combined.


// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
