##### Run Command:

`$ go build command-line-subcommands.go`

___
#### First invoke the foo subcommand.
___
##### Run Command:

`$ ./command-line-subcommands foo -enable -name=joe a1 a2`

##### Results:

`subcommand 'foo'`
`  enable: true`
`  name: joe`
`  tail: [a1 a2]`
___
#### Now try bar.
___
##### Run Command:

`$ ./command-line-subcommands bar -level 8 a1`

##### Results:

`subcommand 'bar'`
`  level: 8`
`  tail: [a1]`
___
#### But bar won't accept foo's flags.
___
##### Run Command:

`$ ./command-line-subcommands bar -enable a1`

##### Results:

`flag provided but not defined: -enable`
`Usage of bar:`
`  -level int`
`        level`
