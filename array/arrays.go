package main

import "fmt"

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

	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}
