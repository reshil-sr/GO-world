package main

import (
	"fmt"
	"time"
)

/*
This is to explain the nil channle and how its useful
ref https://www.youtube.com/watch?v=t9bEg2A4jsw justforfunc #26 is nil channel useful
here we recieve values from two diff chanels and whenever we receive a value we will write that to a third channel
*/

//This is my implementation of the same, its working but the better way is the way compoy explains it
//this only convers the first part of the explanation, this does not cover the issues explained in the second part of the video

func main() {
	chn1 := make(chan int)
	chn2 := make(chan int)
	resChn := make(chan int)

	//sender1
	go func() {
		for i := 1; i < 4; i++ {
			chn1 <- i
		}
		close(chn1)
	}()

	//sender2
	go func() {
		for i := 4; i < 7; i++ {
			chn2 <- i * 2
		}
		close(chn2)
	}()

	go callRec(chn1, chn2, resChn)

	for v := range resChn {
		fmt.Println(v)
	}
}

func callRec(ch1, ch2 <-chan int, resChn chan<- int) {
	defer close(resChn)
	chan1Done, chan2Done := false, false
	for !chan1Done || !chan2Done {
		select {
		case onChn1, ok := <-ch1:
			if !ok {
				chan1Done = true
				continue
			}
			resChn <- onChn1
		case onChn2, ok2 := <-ch2:
			if !ok2 {
				chan2Done = true
				continue
			}
			resChn <- onChn2
		}
		time.Sleep(500 * time.Millisecond)
	}
}
