package main

import "fmt"

//========================================//
//Constants
//#immutable, but can be shadowed
//#replaced by the compiler at compile time
//## value must be calculable at compile time
//#Name like variable
//##pascalCase for exported constants
//## camelCase for internal constants
//# Typed constants work like immutable variables
//##can interoperate only with same type
//#untyped constants work like literals
//## can interoperate with similar types
//#Enumerated constants
//##Special symbol iota allows related constants tobe created easily
//## Iota starts at 0 in each const block and increments by one
//##Watch out of constant values that match zero values for variables
//# Enumerated Expressions
//## operations that can be determined at compile time are allowed
//###Arithmetic
//###Bitwise operations
//###Bitwshifting
//========================================//
const (
	a = iota
	//b = iota
	//c = iota
	b //scope to the constant block
	c //increment the value automatically
)

// const (
// 	a2 = iota //reset value in another block automatically
// )

// const (
// 	//errorSpecialist = iota
// 	_ = iota
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )
// const (
// 	_  = iota //ignore first value by assigning to blank identifier
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrika
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// write your code here
	//const myConst int = 43
	//const fConst float64 = math.Sin(1.57) // constant are not allowed to put value that requires in runtime to set in
	//constant cannot reassain with another valure
	// fmt.Printf("%v, %T\n", a, a)
	//cannot add two types of variable
	// eg const a int =42 and var b int16 = 27
	// but if we declare variable without type it will accept the operations
	//eg const a = 42 and var b int16 = 27 it will take the first variable as a numerical value
	// we can do impulsive operation
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)

	//var specialistType int = catSpecialist
	// var specialistType int
	// fmt.Printf("%v\n", specialistType == catSpecialist)

	// fileSize := 4000000000.
	// fmt.Printf("%.2fGB", fileSize/GB)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

}
