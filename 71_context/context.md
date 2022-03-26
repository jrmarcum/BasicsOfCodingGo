
___
#### To run the program, make sure you are in the program's folder, then use the Run Command below.
___
# Run the server in the background.

___
Run Command:

$ go run context.go

___
# Simulate a client request to `/hello`, hitting
# Ctrl+C shortly after starting to signal
# cancellation.

___
Run Command:

$ curl localhost:8090/hello


Results:

server: hello handler started
^C
server: context canceled
server: hello handler ended

___

###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.