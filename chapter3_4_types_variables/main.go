/*
- Sometimes make a distinction between types and tokens.
- For example suppose you have a dog named Max.
- Max is the token (A particular instance of the member)
- dog is the type (The general concept)
- Go integer type are
	* uint8, uint16, uint32, uint64.
	* int8, int16, int32, int64
- 8,16,32,64 means how many bits each of the types use
- uint means "Unsigned integer" - Only contains positive numbers (or zero)
- int means "Signed integer"
- In additional there two alias types:
	* byte: which is the same as uint8
	* rune: which is the same as int32
- (1 byte = 8 bits, 1024 bytes = 1 kilobyte, 1024 kilobytes = 1 megabyte)

- GO float types are:
	* float32, float64
	* And numbers with imaginary part (complex64 and complex128)

- GO Strings
	* \n get repleaced with a newline
	* \t get repleaced with a tab character
	* len: length of a string
- Go booleans
	* true && true = true
	* true && false = false
	* false && true = false
	* false && false = false
	* true || true = true
	* true || false = true
	* false || true = true
	* false || false = false
- GO variables
	* a variable is a storage
	* var x string
	* x = "first"

	* x := "first"
	* var x = "first"

	* constants are basically variables whose values cannot be changed later.
	* const -> keyword

	* Multiply variables
	* var or const (
		a = 5
		b = 10
		c = 15
	)
*/
package main

import "fmt"

var x string = "first" // Scope of the variable

func main() {
	// fmt.Println("1 + 1 = ", 1 + 1 )
	// fmt.Println("1 + 1 = ", 1.0 + 1.0 )
	// fmt.Println(len("Hello World"))
	// fmt.Println("Hello World"[1])
	// fmt.Println("Hello" + "World")
	// var x string = "Hello World"
	// fmt.Println(x)
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
