#### The spawned programs return output that is the same as if we had run them directly from the command-line. Note: this example uses Unix commands (`date`, `grep`, `bash`) and runs on Unix-like systems.
___
##### Run Command:

`$ go run spawning-processes.go`

##### Results:

`> date`
`Thu 05 May 2022 10:10:12 PM PDT`
`command exit rc = 1`
`> grep hello`
`hello grep`
`> ls -a -l -h`
`drwxr-xr-x  4 mark 136B Oct 3 16:29 .`
`drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..`
`-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go`
