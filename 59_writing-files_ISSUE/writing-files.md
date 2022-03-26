

_______________________________________________________________________________
# Try running the file-writing code.

_______________________________________________________________________________
Run Command:

$ go run writing-files.go

_______________________________________________________________________________
Results:
 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

_______________________________________________________________________________
# Then check the contents of the written files.

_______________________________________________________________________________
Run Command:

$ cat ./tmp/dat1.txt

_______________________________________________________________________________
Results:

hello
go

_______________________________________________________________________________
Run Command:

$ cat ./tmp/dat2.txt

_______________________________________________________________________________
Results:

some
writes
buffered

___

###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.