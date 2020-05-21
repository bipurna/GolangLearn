package main

import (
	"fmt"
)

func main() {
	// write your code here
	const myConst int = 43
	//const fConst float64 = math.Sin(1.57) // constant are not allowed to put value that requires in runtime to set in
	//constant cannot reassain with another valure
	fmt.Printf("%v, %T\n", myConst, myConst)

}
