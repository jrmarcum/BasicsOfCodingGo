/* Go has various variable types. Here are the basic types available.

Invalid: {Invalid, 0, "invalid type"},

Bool:          {Bool, IsBoolean, "bool"},

Int:           {Int, IsInteger, "int"},
Int8:          {Int8, IsInteger, "int8"},
Int16:         {Int16, IsInteger, "int16"},
Int32:         {Int32, IsInteger, "int32"},
Int64:         {Int64, IsInteger, "int64"},

Uint:          {Uint, IsInteger | IsUnsigned, "uint"},
Uint8:         {Uint8, IsInteger | IsUnsigned, "uint8"},
Uint16:        {Uint16, IsInteger | IsUnsigned, "uint16"},
Uint32:        {Uint32, IsInteger | IsUnsigned, "uint32"},
Uint64:        {Uint64, IsInteger | IsUnsigned, "uint64"},
Uintptr:       {Uintptr, IsInteger | IsUnsigned, "uintptr"},

Float32:       {Float32, IsFloat, "float32"},
Float64:       {Float64, IsFloat, "float64"},

Complex64:     {Complex64, IsComplex, "complex64"},
Complex128:    {Complex128, IsComplex, "complex128"},

String:        {String, IsString, "string"},

UnsafePointer: {UnsafePointer, 0, "Pointer"},

UntypedBool:    {UntypedBool, IsBoolean | IsUntyped, "untyped bool"},
UntypedInt:     {UntypedInt, IsInteger | IsUntyped, "untyped int"},
UntypedRune:    {UntypedRune, IsInteger | IsUntyped, "untyped rune"},
UntypedFloat:   {UntypedFloat, IsFloat | IsUntyped, "untyped float"},
UntypedComplex: {UntypedComplex, IsComplex | IsUntyped, "untyped complex"},
UntypedString:  {UntypedString, IsString | IsUntyped, "untyped string"},
UntypedNil:     {UntypedNil, IsUntyped, "untyped nil"},
*/

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
