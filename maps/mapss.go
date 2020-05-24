package main

import (
	"fmt"
	"reflect"
)

// 1.Maps
// a. Collection of value types that are accessed via keys
// b.created via literals or via make function
// c. Members accessed via [key] syntax
// i.myMap["key"]="value"
// d. check for presence with "value, ok" form of result
// e. Multiple assignments refer to same underlying data

// 2.structs
// a. collections of disparate data types that describe a single concept
// b.keyed by named fields
//c.Normally created as types, but anonymous structs are allowed
//d.structs are value types
//e.No inheritance, but can use composition via embedding
//f. tags can be added to struct fields to describe field

// type Doctor struct {
// 	number     int
// 	actorName  string
// 	episodes   []string
// 	companions []string
// }
// type Animal struct {
// 	Name   string
// 	Origin string
// }
// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly   bool
// }

type Animal struct {
	Name   string `required max:"200"`
	Origin string
}

func main() {
	//============arrays=====================//
	// statePopulations := make(map[string]int, 10)
	// statePopulations = map[string]int{
	// 	"Clifornia":   39250017,
	// 	"Texas":       27862596,
	// 	"Florida":     20612439,
	// 	"New York":    19745289,
	// 	"Pennsyvania": 12801539,
	// 	"Illinois":    12801539,
	// 	"Ohio":        11614373,
	// }
	// //fmt.Println(statePopulations)
	// //fmt.Println(statePopulations["New York"])
	// //delete(statePopulations, "Georgia")
	// //pop, ok := statePopulations["Oho"]
	// //fmt.Println(pop, ok)

	// _, ok := statePopulations["Ohio"]
	// fmt.Println(ok)
	// fmt.Println(len(statePopulations))

	//end arrays===========//
	//=======================struct==================//
	// aDoctor := Doctor{
	// 	number:    3,
	// 	actorName: "Jon Pertwee",
	// 	episodes:  []string{},
	// 	companions: []string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	}}

	// // fmt.Println(aDoctor)
	// //fmt.Println(aDoctor.actorName)
	// fmt.Println(aDoctor.companions[1])
	// aDoctor := struct{ name string }{name: "John Pertwee"}
	// anotherDoctor := aDoctor
	// //anotherDoctor := &aDoctor
	// anotherDoctor.name = "Tom Baker"
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)
	//==================================//
	// b := Bird{
	// 	Animal: Animal{Name: "Emu", Origin: "Australia"},

	// 	SpeedKPH: 48,
	// 	CanFly:   false,
	// }
	// // b.Name = "Emu"
	// // b.Origin = "Australia"
	// // b.SpeedKPH = 48
	// // b.CanFly = false

	// fmt.Println(b.Name)
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	//end of struct
}
