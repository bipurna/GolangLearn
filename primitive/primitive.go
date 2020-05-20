package main

import "fmt"

//# Boolean type
////# Values are true or false
////# Not an alias for other types (e.g. int)
////# Zero value is false

//## Numeric types
////#Integers Type
////=>Signed integers
////==> int type has varying size, but min 32 bits
////==> 8bit (int 8) through 64 bit int(64)
////=>Unsigned integers
////==> 8 bit (byte and uint8) through 32 bit (uint32)
////=> Arithmetic operations
////==> addition, subtraction, multiplication, division, remainder
////=>Bitwise operations
////==> and,or , xor,andnot
////=>Zero value is 0
//Can't mix types in same family (uint16 + uint32= error)

////## floating points
////=>floating point numbers
////==> follow IEEE-754 standerd
////==> zero value is 0
////==> 32 and 64 bit versions
////==> literal styles
////====>decimal(3.14)
////====>exponential(13e18 or 2E10)
////====>Mixed(13.7e12)
////==> Arithmetic operations
////====>addition, subtraction, multiplication, division (not remainder)

////=> Complex numbers
////==> zero value is (0+0i)
////==> 64 and 128 bit versions
////==> built in functions
////====>complex -make complex number from two floats
////====>real-get real part as float
////====>imag-get imaginary part as float
////====> addition, subtraction, multiplication, division, remainder

////=>text types
////==> strings
////====>UTF-8
////====> immutable
////====> can be concatenated with plus(+) operator
////====> can be converted to []byte
////==> Rune
////====>UTF-32
////====>Alias for int32
////====>special methods normally required to process
////====> e.g. strings.Reader#ReadRune

func main() {
	// go should not worry about var value to init
	// o value is acceptable

	//ex
	var k bool
	fmt.Printf("%v, %T\n", k, k)
	//end ex

	// //#ex
	// n := 1 == 1
	// m := 1 == 2

	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T", m, m)
	// //end ex

	//uint
	// var n uint16 = 42 //( uint8 => 0-255, uint16=> 0-65535, uint32=> 0-4294967295)
	// fmt.Printf("%v, %T\n", n, n)

	//end uint

	//operator

	//// integer
	// a := 10
	// b := 3

	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)
	/// end integer

	// /// different types of integer not allowed to operate
	// // example
	// var i int = 10
	// var j int8 = 3
	// //fmt.Println(i + j)      //produces this error (invalid operation: i + j (mismatched types int and int8))
	// fmt.Println(i + int(j)) //can implicit data conversion
	// //end example

	//bit

	// i := 10 // 1010
	// j := 3  // 0011

	// fmt.Println(i & j)  //1010 & 0011 => 0010 = 2
	// fmt.Println(i | j)  // 1010 | 0011 => 1011 = 11
	// fmt.Println(i ^ j)  // 1010 | 0011 => 1001 = 9
	// fmt.Println(i &^ j) //1010 | 0011 => 0100 = 8

	//end of bit

	////bit shifing
	// a := 8              // 0100 =>2^3
	// fmt.Println(a << 3) //1000000 => 2^3 * 2^3 = 2^6
	// fmt.Println(a >> 3) //0001    => 2^3 / 2^3 = 2^0
	////end bit shifting

	// //floating number (float64  ±2.23E-308 -±1.8E308)

	// n := 3.14
	// n = 13.7e72
	// n = 2.1E14
	// fmt.Printf("%v, %T", n, n)
	// //end of floating number

	// //floating number arithmatic operation
	// a := 10.2
	// b := 3.7

	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// //not available for remainder operator
	//end float

	// complex64 and complex64 complex128
	// a := 1 + 2i
	// b := 2 + 5.2i
	// fmt.Printf("%v, %T\n", real(a), real(a))
	// fmt.Printf("%v, %T", imag(b), imag(b))

	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	//end of complex

	// string includes all UTF-8 - an array of letters

	// s := "this is a string "
	// fmt.Printf("%v, %T\n", s, s)               //=>this is a string, string
	// fmt.Printf("%v, %T\n", s[2], s[2])         //=> 105, uint8
	// fmt.Printf("%v, %T\n", string(s[2]), s[2]) //=>i, uint8
	// // strings are immutable cannot replace value.

	// // string concate +
	// s2 := "this is also a string"
	// fmt.Printf("%v, %T\n", s+s2, s+s2) //=>this is a string this is also a string, string

	// //string can convert to collection of byte
	// b := []byte(s)

	// fmt.Printf("%v, %T\n", b, b) //=>[116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103 32], []uint8

	// rune is int32
	// r := 'a'
	// var r rune = 'a'
	// fmt.Printf("%v, %T\n", r, r)

}
