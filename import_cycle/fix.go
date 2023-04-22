package main

import (
	"GOWORLD/import_cycle/p1"
)

func main() {
	pp1 := p1.PP1Fix{}
	pp1.HelloFromP2SideFix() // Prints: "Hello from package p2"
}
