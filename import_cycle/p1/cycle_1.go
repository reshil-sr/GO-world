package p1

/* import (
	"GOWORLD/import_cycle/p2" //import cycle not allowed
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
	pp2 := p2.New()
	pp2.HelloFromP2()
} */

//fix this using interface

import (
	"GOWORLD/import_cycle/p2"
	"fmt"
)

type PP1Fix struct{}

func NewFix() *PP1Fix {
	return &PP1Fix{}
}

func (p *PP1Fix) HelloFromP1Fix() {
	fmt.Println("Hello from package p1")
}

func (p *PP1Fix) HelloFromP2SideFix() {
	pp2 := p2.NewFix(p)
	pp2.HelloFromP2Fix()
}
