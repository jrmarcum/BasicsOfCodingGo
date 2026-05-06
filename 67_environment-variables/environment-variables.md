#### Running the program shows that we pick up the value for `FOO` that we set in the program, but that `BAR` is empty.
___
##### Run Command:

`$ go run environment-variables.go`

##### Results:

`FOO: 1`
`BAR: `
`...`
___
#### The total list of keys in the environment will depend on your particular machine.
___
#### If we set `BAR` in the environment first, the running program picks that value up.
___
##### Run Command:

`$ $Env:BAR=2 go run environment-variables.go` (for Windows)

`$ BAR=2 go run environment-variables.go` (for Unix-like environments)

##### Results:

`FOO: 1`
`BAR: 2`
`...`
