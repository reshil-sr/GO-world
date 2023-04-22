package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
	// "log"
)

func main() {
	fmt.Println("Hello World!")
	tg := os.Getenv("APP_ENV")
	fmt.Printf("%v, %T\n", tg, tg)
	/* var tg string
	tg = "hai"
	log.Println("The string is: ", tg)

	changeValueUsingPntr(&tg)
	log.Println("The string after updation is: ", tg) */
	// var arr [3]int
	// arr[0] = 1
	/* arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	sum := 0
	for i, _ := range arr {
		sum += arr[i]
	}
	fmt.Println(sum) */
	//map
	/* ma := map[string]int{"one": 1, "two": 2}
	fmt.Println(ma) */

	/* a, b := mathOperations(21, 7)
	fmt.Println(a, b) */
	//adv.Conc()
	//prntStruct()
	// tryMe()
	//tryChannel()
}

/* func mathOperations(x, y int) (z1, z2 int) {
	defer fmt.Println("calculation done!")
	z1 = x + y
	z2 = x - y
	return
}

func changeValueUsingPntr(s *string) {
	tgg := "hello"
	*s = tgg
} */

func prntStruct() {
	type person struct {
		name string
		age  int
	}
	p := person{"Reshil", 34}
	op := new(person)
	op.name = "REshil A S"
	t := &p.age
	fmt.Println(&p.age, *t, op)
}

func tryMe() {
	person := []string{
		"Reshil",
		"Reshma",
		"sandhya",
	}
	p := person[1:2]
	fmt.Println(cap(p))
	p = append(p, "Kunju")
	fmt.Println("Before", person, p)
	person[1] = "Reshma A S"
	fmt.Println("After", person, p)
	fmt.Println(p[0])
	fmt.Printf("person: %p \n", &person)
	fmt.Printf("person at position 1: %p \n", &person[1])
	fmt.Printf("first: %p \n", &p[0])
	fmt.Printf("second: %p \n", &p[1])
}

func tryChannel() {
	dataChan := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dataChan <- doWork()
			}()
		}
		wg.Wait()
		close(dataChan)
	}()
	// go func() {
	for n := range dataChan {
		fmt.Printf("n=%d\n", n)
	}
	// }()
}

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
