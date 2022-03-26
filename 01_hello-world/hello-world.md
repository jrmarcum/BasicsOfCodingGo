// This work was originally from Mark McGranaghan at
// https://github.com/mmcgrana/gobyexample
// And licensed under a Creative Commons Attribution 3.0 Unported License
// http://creativecommons.org/licenses/by/3.0/
// It has been used to provide an example base for multiple languages to
// provide a basis of comparitive programming language study for syntax,
// language simplicity, number of lines of code and operations required to
// perform the same task, as well as compile and run speed combined.

_______________________________________________________________________________
# To run the program, put the code in `hello-world.go` and
# use `go run`.

_______________________________________________________________________________
Run Command:

$ go run hello-world.go

_______________________________________________________________________________
Results:

hello world

_______________________________________________________________________________
# Sometimes we'll want to build our programs into
# binaries. We can do this using `go build`.

_______________________________________________________________________________
Run Command:

$ go build hello-world.go

$ ls

_______________________________________________________________________________
Results:

hello-world.exe (or hello-world)
hello-world.go
hello-worl.txt

_______________________________________________________________________________
# We can then execute the built binary directly.

_______________________________________________________________________________
Run Command:

$ ./hello-world

_______________________________________________________________________________
Results:

hello world
