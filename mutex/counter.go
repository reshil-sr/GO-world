package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock  sync.Mutex
	count int
)

func main() {
	//cntrWithSlp()
	//cntrWithWg()
	//cntrWithCh()
	/**
		The issue here is that count++ is actually behind the hood is as like below
		temp := count
		temp = temp+1
		count = temp
		so the timing of assigning in step 1 and step 3 may vary and each iteration need not update it properly
		This is the best use case of mutex
		when mutex in place the increment will be atomic ie only one single go routine will able to update the count at a time
		Also we have RWMutex
	**/
	cntrWithMutex()
}

func cntrWithSlp() {
	iter := 1000
	wg := new(sync.WaitGroup)
	wg.Add(iter)
	for i := 0; i < iter; i++ {
		go printCounterSlp()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("the counter is:", count)
}

func printCounterSlp() {
	count++
}

func cntrWithWg() {
	iter := 1000
	wg := new(sync.WaitGroup)
	wg.Add(iter)
	for i := 0; i < iter; i++ {
		go printCounterWg(wg)
	}
	wg.Wait()
	fmt.Println("the counter is:", count)
}
func printCounterWg(wg *sync.WaitGroup) {
	count++
	wg.Done()
}

func cntrWithCh() {
	iter := 1000
	done := make(chan int)
	for i := 0; i < iter; i++ {
		go printCounterCh(done)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("the counter is:", <-done)
	//panic("ss")
}
func printCounterCh(done chan<- int) {
	count++
	done <- count

	//close(done)
}

func cntrWithMutex() {
	iter := 1000
	for i := 0; i < iter; i++ {
		go printCounterMutex()
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println("the counter is:", count)
}
func printCounterMutex() {
	lock.Lock()
	count++
	lock.Unlock()
}
