package main

import (
	"fmt"
	"log"
	"time"
)

// https://medium.com/@func25/what-you-know-about-defer-in-go-is-not-enough-2681d4b128c3
/*
	    • defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns.
			Defer is commonly used to simplify functions that perform various clean-up actions.
        ◦ A deferred function’s arguments are evaluated when the defer statement is evaluated.
        ◦ Deferred function calls are executed in Last In First Out order after the surrounding function returns.
        ◦ Deferred functions may read and assign to the returning function’s named return values.
*/

func main() {
	// a()
	// b()
	//fmt.Println(c())
	checkTimer()
}

func a() {
	//A deferred function’s arguments are evaluated when the defer statement is evaluated.
	i := 0
	//the expression “i” is evaluated when the Println call is deferred
	defer fmt.Println(i) // will return 0 since the value of i at this point of time is 0
	i++
	defer fmt.Println(i) // will return 1 since the value of i at this point of time is 1, its incr in prv line
	return
	//Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	//hence result 1,0
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
	//prints “3210”:
}

func c() (i int) {

	//Deferred functions may read and assign to the returning function’s named return values.
	defer func() {
		fmt.Println("value of i inside defer: ", i)
		i++
	}()
	fmt.Println("value of i: ", i)
	return 2 //this will return 3
}

func checkTimer() {
	fmt.Println("Before")
	defer func() {
		fmt.Println("First Defer : ")
	}()
	//defer timer("Checking Timer")()
	defer func() {
		fmt.Println("second Defer : ")
	}()
	fmt.Println("After")
	return
	//output
	/*
		Before
		In-Before
		In-After, Time: 2023-03-14 21:30:09.662194228 +0530 IST m=+0.000048307
		After
		second Defer :
		2023/03/14 21:30:10 Checking Timer took 1.000336605s
		First Defer :
	*/
}

func timer(name string) func() {
	fmt.Println("In-Before")
	start := time.Now()
	fmt.Println("In-After, Time:", start)
	time.Sleep(1 * time.Second)
	return func() {
		log.Printf("%s took %v\n", name, time.Since(start).String())
	}
}
