

_______________________________________________________________________________
# Note that while slices are different types than arrays,
# they are rendered similarly by `fmt.Println`.

_______________________________________________________________________________
Run Command:

$ go run slices.go

_______________________________________________________________________________
Results:

emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

_______________________________________________________________________________
# Check out this [great blog post](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)
# by the Go team for more details on the design and
# implementation of slices in Go.

___

###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.