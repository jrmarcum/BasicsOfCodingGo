#### Run the server in the background, then simulate a client request to `/hello`. Hit Ctrl+C shortly after to signal cancellation.
___
##### Run Command:

`$ go run context.go &`
`$ curl localhost:8090/hello`

##### Results:

`server: hello handler started`
`^C`
`server: context canceled`
`server: hello handler ended`
