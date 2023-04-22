package p1

/*
The Interface Way:

    Package p1 to use functions/variables from package p2 by importing package p2.
    Package p2 to call functions/variables from package p1 without having to import package p1.
	All it needs to be passed package p1 instances which implement an interface defined in p2, those instances will be views as package p2 object.

That’s how package p2 ignores the existence of package p1 and import cycle is broken.
*/
//main method is in cycle.go
import (
	"GOWORLD/import_cycle/p2"
	"fmt"
)

type PP1 struct{}

func New() *PP1 {
	return &PP1{}
}

func (p *PP1) HelloFromP1() {
	fmt.Println("Hello from package p1")
}

func (p *PP1) HelloFromP2Side() {
	pp2 := p2.New(p)
	pp2.HelloFromP2()
}
