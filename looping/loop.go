package main

import "fmt"

// 1.for statements
// a.simple loops
// i.for initializer; test; incrementer {}
// ii. for test{}
// iii. for {}

// b.exiting early
// i.break
// ii.continue
// iii.labels
// c. looping over collections
// i.arrays, slices,maps,strings,channels
// ii. for k, v := range collection {}
func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	// Loop:
	// 	for i := 1; i <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			fmt.Println(i * j)
	// 			if i*j >= 3 {
	// 				break Loop
	// 			}
	// 		}
	// 	}

	// statePopulations := map[string]int{
	// 	"Clifornia":   39250017,
	// 	"Texas":       27862596,
	// 	"Florida":     20612439,
	// 	"New York":    19745289,
	// 	"Pennsyvania": 12801539,
	// 	"Illinois":    12801539,
	// 	"Ohio":        11614373,
	// }
	// for k, v := range statePopulations {
	// 	fmt.Println(k, v)
	// }

	// s := []int{1, 2, 3}
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }
	s := "hello go!"
	for k, v := range s {
		//fmt.Println(k, v)
		fmt.Println(k, string(v))
	}
}
