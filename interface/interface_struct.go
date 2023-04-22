package main

import "fmt"

type printer interface {
	print() string
}

type myPrinter struct {
	myP1 string
	myP2 int
}

func (mp myPrinter) print() string {
	return "My Printer printing..."
}

func newPrinter() printer {
	return &myPrinter{myP1: "Reshil", myP2: 35}
}

func main() {
	myPr := myPrinter{}
	tg := myPr.print()
	fmt.Println(tg)
	tg2 := newPrinter()
	fmt.Printf("%T\n", tg2)
	fmt.Println(tg2)
	s, ok := tg2.(*myPrinter)
	fmt.Println(s, ok)
}
