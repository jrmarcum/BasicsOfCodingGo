___
#### To run the program, make sure you are in the program's folder, then use the Run Command below.
___
#### Create a new file with some data to read.
___
##### Run Command:

`$ echo "hello" > ./tmp/dat.txt`

`$ echo "go" >> ./tmp/dat.txt`

___
##### Run Command:

`$ go run reading-files.go`

##### Results:
```
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello
```
___
###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.