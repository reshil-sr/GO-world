package main

import (
	"fmt"
	"log"
	"time"
)

/*
This is to explain the nil channle and how its useful
ref https://www.youtube.com/watch?v=t9bEg2A4jsw justforfunc #26 is nil channel useful
here we recieve values from two diff chanels and whenever we receive a value we will write that to a third channel
*/

func main() {
	a := asChan(1, 2, 3, 4)
	b := asChan(5, 6, 7, 8)
	//c := merge(a, b)
	c := mergeWithNIl(a, b)

	for v := range c {
		fmt.Println(v)
	}
}

func asChan(inp ...int) <-chan int {
	chn := make(chan int)
	go func() {
		for _, v := range inp {
			chn <- v
			time.Sleep(500 * time.Millisecond)
		}
		close(chn)
	}()
	return chn
}

/* This will use a lot of cpu this can be viewed from the log
2022/12/25 22:34:42 a is done!
2022/12/25 22:34:42 a is done!
2022/12/25 22:34:42 a is done!
2022/12/25 22:34:42 b is done!
even though we have checked chanel is closed, we are entering there
*/
func merge(a, b <-chan int) <-chan int {
	resChn := make(chan int)
	go func() {
		defer close(resChn)
		aDone, bDone := false, false
		for !aDone || !bDone {
			select {
			case v, ok := <-a:
				if !ok {
					aDone = true
					log.Printf("a is done!")
					continue
				}
				resChn <- v
			case v, ok := <-b:
				if !ok {
					bDone = true
					log.Printf("b is done!")
					continue
				}
				resChn <- v
			}
		}
	}()
	return resChn
}

/*
Here we are making the chanel to nil ie a=nil and b=nil
writing or reading from the nil channel blocks for ever, we are utiizing that property here
To restrict further access to the select case
*/
func mergeWithNIl(a, b <-chan int) <-chan int {
	resChn := make(chan int)
	go func() {
		defer close(resChn)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					log.Printf("a is done!")
					continue
				}
				resChn <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					log.Printf("b is done!")
					continue
				}
				resChn <- v
			}
		}
	}()
	return resChn
}
