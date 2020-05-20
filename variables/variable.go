package main

import (
	"fmt"
	"strconv" //string conversion package
)

//==============================================//
// # variable declaration
//// # var foo int
//// # var foo int =42
//// # foo:=42
// ## can't redeclare variables , but can shadow them
//### All variables must be used
//#### Visibility
////#lower case first letter for package scope
////## upper case first letter to export
////### no private scope
//##### Naming Conventions
////# Pascal or camelCase
////////# Capitalize acronyms(HTTP, URL)
////## As short as reasonable
////////# longer names for longer lives
//###### Type conversion
////#destinationType(variable)
////#use strconv package for strings
////end

//==============================================//

//var actorName string = "Elisabeth Sladen"
//var comapnion string = "Sarah Jane Smith"
//var doctorNumber int = 3
//var season int = 11

// var (
// 	actorName    string = "Elisabeth Sladen"
// 	comapnion    string = "Sarah Jane Smith"
// 	doctorNumber int    = 3
// 	season       int    = 11
// )
//var i int = 27

/// variable declare but not used will produce run time error that will keep go
///clean and tidy

//Acronym variable must be all UPPERCASE
//var theHTTP string = "https://google.com"
func main() {
	//var i int
	//i = 42
	//i = 27
	// var i int = 42
	// j := 42.
	// fmt.Printf("%v, %T \n", j, j)
	// fmt.Printf("%v, %T", i, i)
	//fmt.Println(i)
	var i int = 42
	fmt.Println(i)

	//converting variables
	//var j float64  //new variable for float
	//j = float64(i) //assign to convert to float from int
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

}
