// This work was originally from Mark McGranaghan at (https://github.com/mmcgrana/gobyexample).
// And licensed under a Creative Commons Attribution 3.0 Unported License (http://creativecommons.org/licenses/by/3.0/)
// It has been used to provide an example base for multiple languages to provide a basis of comparitive programming
// language study on sytax and simplicity of the languages as far as number of lines of code and operations required
// to perform the same task, as well as compile and run speed combined.


// The Go standard library comes with excellent support
// for HTTP clients and servers in the `net/http`
// package. In this example we'll use it to issue simple
// HTTP requests.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a server. `http.Get` is a
	// convenient shortcut around creating an `http.Client`
	// object and calling its `Get` method; it uses the
	// `http.DefaultClient` object which has useful default
	// settings.
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
