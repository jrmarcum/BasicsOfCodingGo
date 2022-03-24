// This work was originally from Mark McGranaghan at
// https://github.com/mmcgrana/gobyexample
// And licensed under a Creative Commons Attribution 3.0 Unported License
// http://creativecommons.org/licenses/by/3.0/
// It has been used to provide an example base for multiple languages to
// provide a basis of comparitive programming language study for syntax,
// language simplicity, number of lines of code and operations required to
// perform the same task, as well as compile and run speed combined.

_______________________________________________________________________________
# Running our program shows a list sorted by string
# length, as desired.

_______________________________________________________________________________
Run Command:

$ go run sorting-by-functions.go

_______________________________________________________________________________
Results:

[kiwi peach banana]

_______________________________________________________________________________
# By following this same pattern of creating a custom
# type, implementing the three `Interface` methods on that
# type, and then calling sort.Sort on a collection of that
# custom type, we can sort Go slices by arbitrary
# functions.
