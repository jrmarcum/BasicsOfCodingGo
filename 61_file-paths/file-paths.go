// This work was originally from Mark McGranaghan at (https://github.com/mmcgrana/gobyexample).
// And licensed under a Creative Commons Attribution 3.0 Unported License (http://creativecommons.org/licenses/by/3.0/)
// It has been used to provide an example base for multiple languages to provide a basis of comparitive programming
// language study on sytax and simplicity of the languages as far as number of lines of code and operations required
// to perform the same task, as well as compile and run speed combined.


// The `filepath` package provides functions to parse
// and construct *file paths* in a way that is portable
// between operating systems; `dir/file` on Linux vs.
// `dir\file` on Windows, for example.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` should be used to construct paths in a
	// portable way. It takes any number of arguments
	// and constructs a hierarchical path from them.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// You should always use `Join` instead of
	// concatenating `/`s or `\`s manually. In addition
	// to providing portability, `Join` will also
	// normalize paths by removing superfluous separators
	// and directory changes.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` and `Base` can be used to split a path to the
	// directory and the file. Alternatively, `Split` will
	// return both in the same call.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// We can check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Some file names have extensions following a dot. We
	// can split the extension out of such names with `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// To find the file's name with the extension removed,
	// use `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` finds a relative path between a *base* and a
	// *target*. It returns an error if the target cannot
	// be made relative to base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
