#### When we run this program it will block waiting for a signal. By typing `ctrl-C` (which the terminal shows as `^C`) we can send a `SIGINT` signal, causing the program to print the cancellation cause and then exit.
___
##### Run Command:

`$ go run signals.go`

##### Results:

`awaiting signal`
`^C`
`interrupt signal received`
`exiting`
