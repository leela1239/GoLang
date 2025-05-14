package main

import (
	"fmt"
	"strconv"
)

func main() {

	const defaultAge = 18 //must be assigned when declared
	//cant use short assignment notation := for const
	fmt.Printf("Default age to vote is %d\n", defaultAge)

	var i int = 42
	var f float64 = float64(i) // int → float64
	var b byte = byte(i)       // int → byte

	var str string = "123"
	num, _ := strconv.Atoi(str) // string → int (requires parsing)
	//strconv.Atoi returns int,error | if converted to int then error is nil

	bytes := []byte("Hello")      // string → []byte
	backToString := string(bytes) // []byte → string

	fmt.Printf("The type of num is %T", num)
	fmt.Println(f, b, num, backToString)

}

//DataTypes

// Category			Type(s)
//------------------------------------------------------------------------------------
// Boolean	    	bool
// Text				string
// Integer			int, int8, int16, int32, int64                (-ve , 0 , +ve)
// Unsigned Int		uint, uint8, uint16, uint32, uint64, uintptr  (0,+ve)
// Aliases			byte (uint8), rune (int32)
// Floating Pt		float32, float64
// Complex			complex64, complex128

// Formats :

// Verb 	| Meaning 											| Example Output
//-------------------------------------------------------------------------------------------
// %v 		| Default format 									| 42, true, {Name: John}
// %+v 		| Default format with field names (for structs) 	| {Name:John Age:25}
// %#v 		| Go syntax representation 							| main.Person{Name:"John"}
// %T 		| Type of the value 								| int, string, Person
// %% 		| Literal % sign 									| %
// %d 		| Decimal integer 									| 42
// %s 		| String 											| Hello
// %q 		| Quoted string with escape characters 				| "Hello\nWorld"
// %f 		| Floating-point (decimal) 							| 3.141593
// %t 		| Boolean 											| true, false
// %x 		| Hexadecimal (for ints or byte slices) 			| 2a, 68656c6c6f
// %p 		| Pointer address 									| 0xc00000a0b0
