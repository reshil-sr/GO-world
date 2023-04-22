package main

import (
	"GOWORLD/import_cycle/p1"
)

func main() {
	pp1 := p1.PP1{}
	pp1.HelloFromP2Side() // Prints: "Hello from package p2"
}
