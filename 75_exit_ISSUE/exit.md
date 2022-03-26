// This work was originally from Mark McGranaghan at
// https://github.com/mmcgrana/gobyexample
// And licensed under a Creative Commons Attribution 3.0 Unported License
// http://creativecommons.org/licenses/by/3.0/
// It has been used to provide an example base for multiple languages to
// provide a basis of comparitive programming language study for syntax,
// language simplicity, number of lines of code and operations required to
// perform the same task, as well as compile and run speed combined.

_______________________________________________________________________________
#  If you run `exit.go` using `go run`, the exit
# will be picked up by `go` and printed.

_______________________________________________________________________________
Run Command:

$ go run exit.go

_______________________________________________________________________________
Results:

exit status 3

_______________________________________________________________________________
# By building and executing a binary you can see
# the status in the terminal.

_______________________________________________________________________________
Run Command:

$ go build exit.go

$ ./exit

$ echo $?

_______________________________________________________________________________
Results:

3

_______________________________________________________________________________
# Note that the `!` from our program never got printed.
