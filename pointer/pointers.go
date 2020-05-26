package main

import "fmt"

//Creating pointer
///pointer types use an * as a prefix to type pointed to
//// *int - a pointer to an integer
///use the addressof operator (&) to get address of variable

//dereferencing pointers
/// dereference a pointer by preceding with an astrik(*)
/// complex types (e.g. structs) are automatically dereferenced

//Create pointers to objects
///can use the addressof operator (&) if value type already exists
//// ms :=myStruct{foo:42}
//// p := &ms
///Use addressof operator before initializer
//// &myStruct{foo:42}
///use the new keyword
//// can't initialize fields at the same time

//types with internal pointers
/// all assignment operations in Go are copy operations
/// slices and maps contain internal pointers, so copies point to same underlying data

func main() {
	// var a int = 42
	// var b *int = &a
	// fmt.Println(a, *b)
	// a = 27
	// fmt.Println(a, *b)
	// *b = 14
	// fmt.Println(a, *b)

	//pointer arithmetic

	//a := [3]int{1, 2, 3}
	// b := &a[0]
	// c := &a[1]
	// fmt.Printf("%v %p %p\n", a, b, c)

	// var ms *myStruct
	// //ms = &myStruct{foo: 42}
	// ms = new(myStruct)
	// (*ms).foo = 42
	// fmt.Println((*ms).foo)

	// b := a
	// fmt.Println(a, b)
	// a[1] = 42
	// fmt.Println(a, b)

	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

// type myStruct struct {
// 	foo int
// }
