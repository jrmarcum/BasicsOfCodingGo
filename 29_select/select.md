#### We receive the values `"one"` and then `"two"` as expected. Note that the total execution time is only ~2 seconds since both the 1 and 2 second sleeps execute concurrently.
___
##### Run Command:

`$ go run select.go`

##### Results:

`received one`
`received two`
