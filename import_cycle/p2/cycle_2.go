package p2

/* import (
	"GOWORLD/import_cycle/p1" //could not import GOWORLD/import_cycle/p1 (no required module provides package "GOWORLD/import_cycle/p1"
	"fmt"
)

type PP2 struct{}

func New() *PP2 {
	return &PP2{}
}

func (p *PP2) HelloFromP2() {
	fmt.Println("Hello from package p2")
}

func (p *PP2) HelloFromP1Side() {
	pp1 := p1.New()
	pp1.HelloFromP1()
} */

//fix this using interface
//check fix.go for the main method
import (
	"fmt"
)

type tg interface {
	HelloFromP1Fix()
}

type PP2Fix struct {
	pp1 tg //interface
}

func NewFix(p1 tg) *PP2Fix {
	return &PP2Fix{pp1: p1}
}

func (p *PP2Fix) HelloFromP2Fix() {
	fmt.Println("Hello from package p2")
}

func (p *PP2Fix) HelloFromP1SideFix() {
	p.pp1.HelloFromP1Fix()
}
