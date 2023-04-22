package main

import (
	"fmt"
	"time"
)

func main() {
	var isAlive chan struct{}
	isAlive = make(chan struct{})
	timeOut := 600 * time.Millisecond
	//timeOut := 2 * time.Second

	//write to isAlive channel after specified time
	go makeChannelAlive(isAlive, timeOut)

	checkIsActive := checkConnectionISActive(isAlive)
	if !checkIsActive {
		fmt.Println("failed")
	} else {
		fmt.Println("success")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("I am done, exiting after 3 seconds!")
}

func checkConnectionISActive(isAlive chan struct{}) bool {
	ticker := time.NewTicker(600 * time.Millisecond)
	timer := time.NewTimer(1400 * time.Millisecond)
	for {
		select {
		case t := <-isAlive:
			fmt.Printf("case isAlive: return %v\n", t)
			return true
		case <-timer.C:
			//case <-time.After(1500 * time.Millisecond):
			fmt.Println("case timer: return false after 1000 milliseconds")
			return false
		case <-ticker.C:
			fmt.Println("case ticker: that ticks on every 600 ms")
		}
	}
}

func makeChannelAlive(ch chan struct{}, timeOut time.Duration) {
	time.Sleep(timeOut)
	ch <- struct{}{}
	fmt.Printf("waited for %v and write to channel\n", timeOut)
}
