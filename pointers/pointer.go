package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")
	//tryDiffThings()
	a := 4
	fmt.Println("address of a: ", &a)
	square(a)
	squareAdd(&a)
}

func tryDiffThings() {
	x, y := 1, 2                       // var x int = 1, var y int = 2
	fmt.Printf("x,y is %v,%v\n", x, y) // x is 1 , y is 2
	fmt.Printf("%T,%T\n", &x, &y)      // *int,*int
	pntr := &x                         // var pntr *int = &x
	fmt.Printf("%T\n", pntr)           // *int
	fmt.Println("pntr Add:", &pntr)    // address where pntr is stored
	fmt.Println(&x, ",", &y)           //address where the x var,y var is stored
	fmt.Println(*pntr)                 //dereferncing oprator * Take the pointer and figure out the value of this.
	pntr = &y                          // var pntr *int = &y, this is possible bcs type of pntr is pointer with the base type integer and y is an integer
	fmt.Println(*pntr)                 //prints 2
	*pntr = 3                          //will change the value at &y
	fmt.Printf("x,y is %v,%v\n", x, y) // x is 1 , y is 3
	z := "hello pointer"
	fmt.Printf("%T\n", &z) // *string,
	//fmt.Println(*&z)       //will act as dereferncing oprt, evn though its shows warning in vscode
}

func square(num int) {
	fmt.Println("Square")
	fmt.Printf("%v, %v\n", &num, num)
	num *= num
	fmt.Printf("%v, %v\n", &num, num)
}

func squareAdd(num *int) {
	fmt.Println("Square Address")
	fmt.Printf("%v, %v\n", num, *num)
	*num *= *num
	fmt.Printf("%v, %v\n", num, *num)
}
