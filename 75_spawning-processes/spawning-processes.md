___
#### To run the program, make sure you are in the program's folder, then use the Run Command below.
___
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

___

###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.
