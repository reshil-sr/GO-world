package adv

import (
	"fmt"
	"sync"
)

func WtGrp() {
	fmt.Println("Welcome to wait group world")
	//step 1: Wait group can be think of as a counter, here we say wait for 1 goroutine to finish execution
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wtCount("sheep")
		wg.Done() // decrement the counter and make it 0
	}()

	wg.Wait() //blocker code wait until the counter is 0
}

func wtCount(s string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(s)
	}
}
