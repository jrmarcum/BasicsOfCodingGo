#### Run the TCP server in the background, then send data and capture the response using netcat.
___
##### Run Command:

`$ go run tcp-server.go &`
`$ echo "Hello from netcat" | nc localhost 8090`

##### Results:

`ACK: HELLO FROM NETCAT`
