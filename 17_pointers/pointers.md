#### `zeroval` doesn't change the `i` in `main`, but `zeroptr` does because it has a reference to the memory address for that variable.
___
##### Run Command:

`$ go run pointers.go`

##### Results:

`initial: 1`
`zeroval: 1`
`zeroptr: 0`
`pointer: 0xc0000b4008`
___
#### Note: The pointer address on the last line changes each time the program is run.
