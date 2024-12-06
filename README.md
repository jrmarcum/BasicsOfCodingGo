# Basics of Coding Go
## Preface
What this text is and what it is not: This text is intended to introduce the reader to the basics of the Go programming language in the sense that they will be able to write minimal types of programs and run the code. It is not intended to go into advanced topics like concurrency and use of pointers, advanced data structures, object oriented programming, testing and debugging techniques and other software engineering principals.

The programs are intended to be run in the terminal as that is common to most operating systems. Linux and Mac come preinstalled with a terminal. Windows may or may not have it pre-installed. "Windows Terminal" can be installed from the Microsoft store. Just do a search for it and install it.
## Installation and Setup of Go
To setup Go for use, navigate to the [golang](https://go.dev/doc/install) website and download the appropriate installer for your operating system. Install and follow the directions on how to perform further settings for use. A package manager makes the install process much easier (brew for Mac, chocolatety for Windows, and varies based on linux distrobution). If the language has been installed properly you will be able to type the following command in the terminal command line and receive the response shown:
```
$ go version
go version 1.23.3 windows/amd64
(Note: the version shown here varies with your installed version, and the os and cpu will match your os and cpu type)
```
## Comments
Comments are used to document what your code does so that others can understand it when reviewing your code. Comments also document items that are non-performant in order to debug the code items at a later date.
## Keywords
|           |        |         |             |          |
|:---------:|:------:|:-------:|:-----------:|:--------:|
| break     | case   | chan    | const       | continue |
| default   | defer  | else    | fallthrough | for      |
| func      | go     | goto    | if          | import   |
| interface | map    | package | range       | return   |
| select    | struct | switch  | type        | var      |
## Identifiers

## Operators
> ### Arithmetic Operators
> |        |                                                    |
> |:------:|:---------------------------------------------------|
> | **+**  | add one number to another number                   |
> | **-**  | subtract one number from another number            |
> | **\***  | multiply one number by another number              |
> | **/**  | divide one number by another number                |
> | **%**  | remainder of dividing one number by another number |
> | **++** | increase the value of a variable by 1              |
> | **--** | decrease the value of a variable by 1              |
> ### Comparison Operators
> |        |                                                         |
> |:------:|:--------------------------------------------------------|
> | **==** | check if a variable is equal to another                 |
> | **!=** | check if a variable is not equal to another             |
> | **>**  | check if a varaible is greater than another             |
> | **<**  | check if a variable is less than another                |
> | **>=** | check if a variable is greater than or equal to another |
> | **<=** | check if a variable is less than or equal to another    |
> ### Logical Operators
> |        |                                                 |
> |:------:|:------------------------------------------------|
> | **&&** | returns true if both statements are true        |
> | **\|\|** | returns true if one of the statements are true  |
> | **!**  | reverse the result of a true or false statement |
> ### Assignment Operators
> |         |                                                                                                              |
> |:-------:|:-------------------------------------------------------------------------------------------------------------|
> | **=**   | assign a data type to a variable                                                                             |
> | **+=**  | add a number to the existing value of a variable and assign the result to variable                           |
> | **-=**  | subtract a number to the existing value of a variable and assign the result to variable                      |
> | **\*=** | multuply a number to the existing value of a variable and assign the result to variable                      |
> | **/=**  | divide the existing value of a variable by a number and assign the result to variable                        |
> | **%=**  | take the remainder of the existing value of a variable divided by a number and assign the result to variable |
## Data Types
  ###  1.  Variables
>> #### a.  **String, string, str**: used for storing text and/or characters
>> 
>> #### b.  **Char, char, Character, character**: a single character/letter/number, or ASCII values, UTF-8 code unit
>> 
>> #### c.  **wchar**: UTF-16 code unit
>> 
>> #### d.  **dchar**: UTF-32 code unit and Unicode code point
>> 
>> #### e.  **Numbers**
>>
>>> #####  1)  *Number, number, numeric*: stores numeric data with or without decimal
>>>
>>> #####  2)  *Int8, int8, i8,sbyte*: stores positive or negative integers from -2^7 to (2^7)-1, with 3 significant figure precision
>>>
>>> #####  3)  *Int16, int16, i16, Short, short*: stores positive or negative integers from -2^15 to (2^15)-1, with 5 significant figure precision
>>>
>>> #####  4)  *Int32, int32, i32, Int, int, Integer, integer*: stores positive or negative integers from -2^31 to (2^31)-1, with 10 significant figure precision
>>>
>>> #####  5)  *Int64, int64, i64, bigint, Long, long*: stores positive or negative integers from -2^63 to (2^63)-1, with 19 significant figure precision
>>>
>>> #####  6)  *Int128, i128*: stores positive or negative integers from -2^127 to (2^127)-1, with 39 significant figure precision
>>>
>>> #####  7)  *isize*: same as i32 or i64 depending on computer architecture
>>>
>>> #####  8)  UInt8, unint8, u8, ubyte, Byte, byte, bytes*: stores positive integers from 0 to (2^8)-1, with 3 significant figure precision
>>>
>>> #####  9)  *UInt16, unint16, u16, ushort*: stores positive integers from 0 to (2^16)-1, with 5 significant figure precision
>>>
>>> ##### 10)  *UInt32, uint32, unint32, u32, uint*: stores positive integers from 0 to (2^32)-1, with 10 significant figure precision
>>>
>>> ##### 11)  *UInt64, uint64, unint64, u64, ulong*: stores positive integers from 0 to (2^64)-1, with 19 significant figure precision
>>>
>>> ##### 12)  *UInt128, u128*: stores positive integer numbers from 0 to (2^128)-1, with 39 significant figure precision
>>>
>>> ##### 13)  *usize*: same as u32 or u64 depending on computer architecture
>>>
>>> ##### 14)  *Real, real*: either the largest floating point type that the hardware supports, or double; whichever is larger
>>>
>>> ##### 15)  *Float16*: stores fractional numbers from -2^15 to (2^15)-1, with 5 significant figure precision
>>>
>>> ##### 16)  *Float32, float32, f32, Float, float*: stores fractional numbers from -2^31 to (2^31)-1, with 10 significant figure precision 
>>>
>>> ##### 17)  *Double, double, Float64, float64, f64*: stores fractional numbers from -2^63 to (2^63)-1, with 19 significant figure precision
>>>
>>> ##### 18)  *decimal*: stores numbers from -7.9E-28 to +7.9E28 (28 digits of precision)
>>>
>>> ##### 19)  *BigRational*: construction of number from an i32 numerator and i32 denominator
>>>
>>> ##### 20)  *Complex, complex, cfloat*: complex number type made of two floats
>>>
>>> ##### 21)  *cdouble*: complex number type made of two doubles
>>>
>>> ##### 22)  *creal*: complex number type made of two reals 
>>>
>>> ##### 23)  *ifloat*: imaginary value type of float
>>>
>>> ##### 24)  *idouble*: imaginary value type of double
>>>
>>> ##### 25)  *ireal*: imaginary value type of real
> ###  2. **Constants** : 
> ###  3. **Boolean, boolean, Bool, bool, Logical, logical**: values of True, true, False, false or None
> ###  4. **Lists**
> ###  5. **Arrays**      
> ###  6. **Tuples**
> ###  7. **Dictionaries**
> ###  8. **Sets**
> ###  9. **Frozen Set**
## Functions
> ###  1. **Range**: Data range function typically using in for loops
> ###  2. **Date and Time**: produces the date and time in a specified format
## Statements
A statement is an instruction that a program can execute. They are usually made line by line in your coding file. Lines in your code can have multiple statements that are separated typically by semicolons.
> ### Conditional Statements
> |               |                                                                                                      |
> |:-------------:|:-----------------------------------------------------------------------------------------------------|
> |  **if-else**  | performs a statement 'if' a condition is met and if not (ie. 'else') performs the default statement  |
> |  **if-else**  | performs a statement 'if' a condition is met and if not, performs a statement 'else if' a condition\
>                   is met and if not (ie. 'else') performs the default statement                                        |
> ### Iterative Statements
> |                   |                                                                                              |
> |:-----------------:|:---------------------------------------------------------------------------------------------|
> | **for**           | loop statement over a defined range of predetermined variables or values                     |
> | **for-in**        | loop statement over each element in an array                                                 |
> | **for-each**      | loop statement over each element in an array                                                 |
> | **for-of**        | loop statement over each element in an array                                                 |
> | **while**         | indefinite loop that is terminated while a condition is true                                 |
> | **while-let**     | while let destructures a variable into another variable perform a block of code else break   |
> | **if-let**        | if let destructures a variable into another variable perform a block of code                 |
> | **do**            | repeats a bock of code while a boolean condition is true or until the condition becomes true |
> | **do-while**      | repeats a bock of code while a boolean condition is true or until the condition becomes true |            
> | **do-until**      | repeats a bock of code while a boolean condition is true or until the condition becomes true |
> | **do-loop-while** | repeats a bock of code while a boolean condition is true or until the condition becomes true |
> | **do-loop-until** | repeats a bock of code while a boolean condition is true or until the condition becomes true |
> | **repeat-while**  | repeats a bock of code while a boolean condition is true or until the condition becomes true |
> | **repeat-until**  | repeats a bock of code while a boolean condition is true or until the condition becomes true |
> | **repeat-if**     | repeats a bock of code while a boolean condition is true or until the condition becomes true |
> | **loop**          | indefinite loop that requires a break statement to terminate the loop                        |
> ### Transfer Statements
> |         |              |
> |:-------:|:-------------|
> | **break** | terminates a statement at the point of call |
> | **continue** | continues to the next loop increment |
> | **pass** | it is a null statement used to pass the statement flow on to the next statement in a function or class in which we have not determined the correct code to put in it yet |
> | **goto** | sends the control flow to a specific statement |
> | **fallthrough** | in swift this statement forces the program to check each case and not stop at the first valid case |
> ### Switch Statements
> |                          |                                                                                                     |
> |:------------------------:|:----------------------------------------------------------------------------------------------------|
> | Express Switch Statement | uses an evaluated expression to brach to a case that matches the result of the evaluated expression |
> | Type Switch Statement    | uses the type of a variable to branch to a case that matches the variables type                     |
## File Input and Output
> ###  1. File Input
>> #### a. Input casting.
>>
>> #### b. Handling Errors from incorrect input types.
>>
>> #### c. json
>>
>> #### d. html
>>
>> #### e. xml
>>
>> #### f. csv
>>
> ###  2. File Output
>> #### a. String Formatting
>>
>> #### b. json
>>
>> #### c. html
>>
>> #### d. xml
>>
>> #### e. csv
