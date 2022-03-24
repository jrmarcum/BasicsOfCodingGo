// This work was originally from Mark McGranaghan at
// https://github.com/mmcgrana/gobyexample
// And licensed under a Creative Commons Attribution 3.0 Unported License
// http://creativecommons.org/licenses/by/3.0/
// It has been used to provide an example base for multiple languages to
// provide a basis of comparitive programming language study for syntax,
// language simplicity, number of lines of code and operations required to
// perform the same task, as well as compile and run speed combined.

_______________________________________________________________________________
# Running this program will cause it to panic, print
# an error message and goroutine traces, and exit with
# a non-zero status.

_______________________________________________________________________________
Run Command:

$ go run panic.go

_______________________________________________________________________________
Results:

panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:22 +0x45
...
exit status 2

_______________________________________________________________________________
# Note that unlike some languages which use exceptions
# for handling of many errors, in Go it is idiomatic
# to use error-indicating return values wherever possible.
