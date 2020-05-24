package main

import (
	"fmt"
	"math"
)

// 1.if statements
//  a.initializer
//  b.comparison operators
//  c.logical operators
//  d.short circuiting
//  e.if-else statements
//  f.if -else if statements
//  g.equality and floats

// switch statements
//  a.switching on a tag
//  b.cases with multiple test
//  c. initializers
//  d.switches with no tag
//  e.fallthrough
//  f.type switches
//  g. breaking out early

func main() {
	numbers := 50
	guess := 120
	// if guess < numbers {
	// 	fmt.Println("Too Low")
	// }
	// if guess > numbers {
	// 	fmt.Println("Too High")
	// }
	// if guess == numbers {
	// 	fmt.Println("You Got it.")
	// }
	// fmt.Println(numbers <= guess, numbers >= guess, numbers != guess)
	switchStat()
	mcheck()
	if guess < 1 {
		fmt.Println("Guess must be greater than 1 ")
	} else if guess > 100 {
		fmt.Println("Guess must be less than 100")
	} else {
		if guess < numbers {
			fmt.Println("Too Low")
		}
		if guess > numbers {
			fmt.Println("Too High")
		}
		if guess == numbers {
			fmt.Println(" You got it.")
		}
		fmt.Println(numbers <= guess, numbers >= guess, numbers != guess)
	}
}

// func returnTrue() bool {
// 	fmt.Println("returning true")
// 	return true
// }

func mcheck() {
	myNum := 0.123456789
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These aer different")
	}
}

func switchStat() {
	// switch 2 {
	// case 1, 5, 10:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("not both")
	// }

	// i := 10
	// switch {
	// case i <= 10:
	// 	fmt.Println("less than or equal to 10")
	// 	//fallthrough
	// case i <= 20:
	// 	fmt.Println("less than or equal to twenty")
	// default:
	// 	fmt.Println("greater than twenty")
	// }

	// var i interface{} = "1.11" //1//1.1//[3]int{}
	// switch i.(type) {
	// case int, int8, int16, int32, int64:
	// 	fmt.Println("i is int")
	// case float32, float64:
	// 	fmt.Println("i is float")
	// case string:
	// 	fmt.Println("i is string")
	// default:
	// 	fmt.Println("i is another type")
	// }

}
