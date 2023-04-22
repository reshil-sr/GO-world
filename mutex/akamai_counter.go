package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	a_lock  sync.Mutex
	a_count int
	rate    int
)

func main() {
	a_cntrWithMutex()
}

func a_cntrWithMutex() {
	a_count = 10000
	rate = 150
	iter := 10000 / rate
	ticker := time.NewTicker(40 * time.Millisecond)
	//tickerWait := time.NewTicker(5 * time.Second)
	tickerAdd := time.NewTicker(1 * time.Second)
	for i := 0; i < iter; i++ {
		//go processRequest()
		select {
		case <-ticker.C:
			{
				fmt.Printf("it is %d th iteration, another 100 Millisecond & the token remainig is %d\n-> ", i, a_count)
				takeToken()
			}
		case <-tickerAdd.C:
			{
				addToken()
				fmt.Printf("it is %d th iteration, another Second & the token remainig is %d\n-> ", i, a_count)
			}
			/* case <-tickerWait.C:
			{
				time.Sleep(5 * time.Second)
				fmt.Printf("it is %d th iteration, Lets wait for 5 second & the token remainig is %d\n-> ", i, a_count)
			} */
		}
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println("the counter is:", a_count)
}
func takeToken() {
	a_lock.Lock()
	a_count = a_count - rate
	a_lock.Unlock()
}

func addToken() {
	a_lock.Lock()
	a_count = a_count + 200
	a_lock.Unlock()
}
