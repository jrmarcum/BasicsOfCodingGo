___
#### To run the program, make sure you are in the program's folder, then use the Run Command below.
___
#### `zeroval` doesn't change the `i` in `main`, but `zeroptr` does because it has a reference to the memory address for that variable.
___
##### Run Command:

`$ go run pointers.go`


##### Results (pointer will vary per run):
```
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100     (Note: this value changes each time the program is run)
```
___

###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.