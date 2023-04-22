package p2

/*
The Interface Way:

    Package p1 to use functions/variables from package p2 by importing package p2.
    Package p2 to call functions/variables from package p1 without having to import package p1.
	All it needs to be passed package p1 instances which implement an interface defined in p2, those instances will be views as package p2 object.

That’s how package p2 ignores the existence of package p1 and import cycle is broken.
*/

import (
	"fmt"
)

type pp1 interface {
	HelloFromP1()
}

type PP2 struct {
	PP1 pp1
}

func New(pp1 pp1) *PP2 {
	return &PP2{
		PP1: pp1,
	}
}

func (p *PP2) HelloFromP2() {
	fmt.Println("Hello from package p2")
}

func (p *PP2) HelloFromP1Side() {
	p.PP1.HelloFromP1()
}
