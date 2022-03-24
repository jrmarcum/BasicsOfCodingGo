// This work was originally from Mark McGranaghan at (https://github.com/mmcgrana/gobyexample).
// And licensed under a Creative Commons Attribution 3.0 Unported License (http://creativecommons.org/licenses/by/3.0/)
// It has been used to provide an example base for multiple languages to provide a basis of comparitive programming
// language study on sytax and simplicity of the languages as far as number of lines of code and operations required
// to perform the same task, as well as compile and run speed combined.


// [Environment variables](http://en.wikipedia.org/wiki/Environment_variable)
// are a universal mechanism for [conveying configuration
// information to Unix programs](http://www.12factor.net/config).
// Let's look at how to set, get, and list environment variables.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.SplitN` them to
	// get the key and value. Here we print all the keys.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
