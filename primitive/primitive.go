package main

import "fmt"

//# Boolean type
//## Numeric types
////# integers
////## floating points
////### complex numbers
//### text types
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

	//floating number (float64  ±2.23E-308 -±1.8E308)

	n := 3.14
	n = 13.7e72
	n = 2.1E14
	fmt.Printf("%v, %T", n, n)
}
