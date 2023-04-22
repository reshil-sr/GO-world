package main

import "fmt"

type A struct {
}

func (a *A) Func1(i int) {
	fmt.Println("called A.Func1")
	a.Func2(i)
}

func (a *A) Func2(i int) {
	fmt.Println("called A.Func2")
	fmt.Println(i)
}

type B struct {
	*A
}

func (b *B) Func2(i int) {
	fmt.Println("called B.Func2")
	i += 1
	b.A.Func2(i)
}

func main() {
	var b = B{}
	fmt.Printf("%#v\n", b.Func2)
	b.Func1(1)
	b.Func2(10)
}
