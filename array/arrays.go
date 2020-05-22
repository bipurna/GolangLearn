package main

import "fmt"

//#Arrays
//##Collection of items with same type
//## Fixed Size
//## Declaration styles
//### a := [3]int{1,2,3}
//### a := [...]int{1,2,3}
//### var a [3] int

//## Access via zero vased index
//### a := [3]int {1,3,5} //a[1]==3

//## len function returns size of array
//## Copies refer to different underlying data

//# Slices
//##Backed by array
//## Creation styles
//### slice existing array or slice
//### literal style
//### via make function
//#### a := make([]int, 10) // create slice with capacity and length ==10
//#### a := make([]int, 100) // create slice with capacity==100 and length ==10

//## len function returns size of array
//## cap function returns length of underlying array
//## append function to add elements to slice
//###May cause expensive copy operation if underlying array is too small

//## Copies refer to same underlying array

func main() {
	//grades := [3]int{97, 85, 93}
	//or
	//grades := [...]int{97, 85, 93}
	// var students [3]string
	// grade1 := 97
	// grade2 := 85
	// grade3 := 93

	// fmt.Printf("Grades: %v\n", grades)
	// fmt.Printf("Grades: %v\n", students)

	// students[0] = "Lisa"
	// students[1] = "Ahmed"
	// students[2] = "Arnold"
	// fmt.Printf("Students: %v, Length: %v\n", students, len(students))

	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}
	// fmt.Println(identityMatrix)

	//a := [...]int{1, 2, 3}
	// a := []int{1, 2, 3}
	// b := a
	// b := &a //pointer to same data
	// b[1] = 5
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Println(b)

	//a := make([]int, 3, 100)
	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1)

	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 2, 3, 4, 5, 6)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	a := []int{1, 2, 3, 4, 5}
	//b := a[:len(a)-1]
	//fmt.Println(b)

	//slicing middle
	b := append(a[:2], a[3:]...)

	fmt.Println(b)

}
