package main

import "fmt"

func main() {
	/**
	For why close the error channel: ?
	Because if you don't close the error channel, the last print statement maybe blocked forever if there is no value in the chan.
	Receive from a closed chan will return immediately even if it's empty.
	**/
	errChn := make(chan error)
	close(errChn)
	fmt.Println(<-errChn) //will print <nil> will be the zero value of your type here since it's error its <nil>
	fmt.Println("end of errChn")

	/*
		Lets say if its an int channel
	*/
	intChn := make(chan int)
	close(intChn)
	vl, open := <-intChn
	fmt.Println(vl, open) //will print 0 and false bcs no more data or returned value is not valid
	//fmt.Printf("%v, %#T\n", <-intChn, <-intChn) // will print 0,int
}
