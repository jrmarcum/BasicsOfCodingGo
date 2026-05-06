#### Running our program shows a list sorted by string length, as desired.
___
##### Run Command:

`$ go run sorting-by-functions.go`

##### Results:

`[kiwi peach banana]`
___
#### By following this same pattern of creating a custom type, implementing the three `Interface` methods on that type, and then calling sort.Sort on a collection of that custom type, we can sort Go slices by arbitrary functions.
