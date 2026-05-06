#### If you run `exit.go` using `go run`, the exit will be picked up by `go` and printed. Note that the `!` from our program never gets printed because `defer` does not run when using `os.Exit`.
___
##### Run Command:

`$ go run exit.go`

##### Results:

`exit status 3`
