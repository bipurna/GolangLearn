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
	var n uint16 = 42 //( uint8 => 0-255, uint16=> 0-65535, uint32=> 0-4294967295)
	fmt.Printf("%v, %T\n", n, n)
}
