#### Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.
___
##### Run Command:

`$ go run panic.go`

##### Results:

`panic: a problem`

`goroutine 1 [running]:`
`main.main()`
`        /.../panic.go:22 +0x45`
`exit status 2`
___
#### Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
