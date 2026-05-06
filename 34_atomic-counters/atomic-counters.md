#### We expect to get exactly 50,000 operations. Had we used a non-atomic integer and incremented it with `ops++`, we'd likely get a different number between runs because the goroutines would interfere with each other.
___
##### Run Command:

`$ go run atomic-counters.go`

##### Results:

`ops: 50000`
