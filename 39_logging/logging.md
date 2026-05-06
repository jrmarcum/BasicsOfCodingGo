___
#### To run the program, make sure you are in the program's folder, then use the Run Command below.
___
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

___

###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.
