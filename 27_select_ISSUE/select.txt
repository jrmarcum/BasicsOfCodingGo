// This work was originally from Mark McGranaghan at
// https://github.com/mmcgrana/gobyexample
// And licensed under a Creative Commons Attribution 3.0 Unported License
// http://creativecommons.org/licenses/by/3.0/
// It has been used to provide an example base for multiple languages to
// provide a basis of comparitive programming language study for syntax,
// language simplicity, number of lines of code and operations required to
// perform the same task, as well as compile and run speed combined.

_______________________________________________________________________________
# We receive the values `"one"` and then `"two"` as
# expected.

_______________________________________________________________________________
Run Command:

$ time go run select.go

_______________________________________________________________________________
Results:

received one
received two

_______________________________________________________________________________
# Note that the total execution time is only ~2 seconds
# since both the 1 and 2 second `Sleeps` execute
# concurrently. real	0m2.245s
