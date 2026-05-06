#### The string encodes to slightly different values with the standard and URL base64 encoders (trailing `+` vs `-`) but they both decode to the original string as desired.
___
##### Run Command:

`$ go run base64-encoding.go`

##### Results:

`YWJjMTIzIT8kKiYoKSctPUB+`
`abc123!?$*&()'-=@~`

`YWJjMTIzIT8kKiYoKSctPUB-`
`abc123!?$*&()'-=@~`
