package main

import (
	"fmt"
	"sync"
)

func main() {
	var numMemPieces int
	numOfWorkers := 1024 * 1024
	wg := new(sync.WaitGroup)
	wg.Add(numOfWorkers)
	memPool := &sync.Pool{
		New: func() interface{} {
			numMemPieces++
			mem := make([]byte, 1024)
			return &mem
		},
	}
	for i := 0; i < numOfWorkers; i++ {
		//go callJobs(wg)
		go func() {
			mem := memPool.Get().(*[]byte)
			//fmt.Printf("%T\n", mem)
			memPool.Put(mem)
			//fmt.Println("wait!")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Number of mem pieces allocated :", numMemPieces)
}

func callJobs(wg *sync.WaitGroup) {
	fmt.Println("entered")
	wg.Done()
}
