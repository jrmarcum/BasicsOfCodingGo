// This work was originally from Mark McGranaghan at
// https://github.com/mmcgrana/gobyexample
// And licensed under a Creative Commons Attribution 3.0 Unported License
// http://creativecommons.org/licenses/by/3.0/
// It has been used to provide an example base for multiple languages to
// provide a basis of comparitive programming language study for syntax,
// language simplicity, number of lines of code and operations required to
// perform the same task, as well as compile and run speed combined.

_______________________________________________________________________________
# Run the server in the background.

_______________________________________________________________________________
Run Command:

$ go run context.go

_______________________________________________________________________________
# Simulate a client request to `/hello`, hitting
# Ctrl+C shortly after starting to signal
# cancellation.

_______________________________________________________________________________
Run Command:

$ curl localhost:8090/hello

_______________________________________________________________________________
Results:

server: hello handler started
^C
server: context canceled
server: hello handler ended
