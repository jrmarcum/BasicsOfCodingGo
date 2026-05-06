#### Sample output; the date and time emitted will depend on when the example ran.
___
##### Run Command:

`$ go run logging.go`

##### Results:

`2023/08/22 10:45:16 standard logger`
`2023/08/22 10:45:16.904141 with micro`
`2023/08/22 10:45:16 logging.go:40: with file/line`
`my:2023/08/22 10:45:16 from mylog`
`ohmy:2023/08/22 10:45:16 from mylog`
`from buflog:buf:2023/08/22 10:45:16 hello`
`{"time":"2023-08-22T10:45:16.904166391-07:00","level":"INFO","msg":"hi there"}`
`{"time":"2023-08-22T10:45:16.904178985-07:00","level":"INFO","msg":"hello again","key":"val","age":25}`
