#### Note that while slices are different types than arrays, they are rendered similarly by `fmt.Println`.
___
##### Run Command:

`$ go run slices.go`

##### Results:

`emp: [  ]`
`set: [a b c]`
`get: c`
`len: 3`
`apd: [a b c d e f]`
`cpy: [a b c d e f]`
`sl1: [c d e]`
`sl2: [a b c d e]`
`sl3: [c d e f]`
`dcl: [g h i]`
`2d:  [[0] [1 2] [2 3 4]]`
___
#### Check out this [great blog post](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html) by the Go team for more details on the design and implementation of slices in Go.
