package main

import (
	"fmt"
	"time"
)

func main() {
	connIsAlive := make(chan struct{}, 1)
	//timeOut := 500 * time.Millisecond
	ticker := time.NewTicker(500 * time.Millisecond)
	ticker2 := time.NewTicker(2 * time.Second)
	timer := time.NewTimer(750 * time.Millisecond)
	//defer ticker2.Stop()
	//defer ticker.Stop()

	for {
		select {
		case t := <-connIsAlive:
			fmt.Printf("Type of connIsAlive is %T\n", t)
			fmt.Printf("Value of connIsAlive is %v\n", t)
		case <-ticker.C:
			//connIsAlive <- struct{}{}
			fmt.Println("Its another 500 ms!")
		case <-ticker2.C:
			fmt.Println("Its now 2s! pls exit")
			return
			//break
		case <-timer.C:
			fmt.Println("Its timer now and its 750ms")
		}
	}
	fmt.Println("I am done")
}
