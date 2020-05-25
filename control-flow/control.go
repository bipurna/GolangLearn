package main

import (
	"fmt"
	"log"
)

// 1.defer
// a. used to delay execution of a statemnet until function exits
// b.useful to group "open" and "close" functions together
//  i. be careful in loops
// c.run in LIFO(last in, first out) order
// d. Arguements evaluated at time defer is executed not at time of called function excution
// 2.panic
// a.occur when program cannot continue at all
// i. dont use when file can't be opened unless it is critical
// ii. Use for unrecoverable events cannot obtain TCP port for web server
// b. function will stop executing
// i. deferred functions will still fire
// c. if nothing handels panic program will exit

// 3.Recover
// a.used to recover from panics
// b.only useful in deferred functions

func main() {
	//==================defer============
	// fmt.Println("start")
	// fmt.Println("middle")
	// fmt.Println("end")
	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
	// a := "start"
	// defer fmt.Println(a)
	// a = "end"
	//==============dd=======
	//===========panicking===========

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// defer fmt.Println("this was deferred")
	// panic("something bad happend")
	// fmt.Println("end")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()
	// panic("something bad happend")
	// fmt.Println("end")

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error:%s", err)
			panic(err)
		}
	}()
	panic("Something bad happend")
	fmt.Println("done panicking")
}
