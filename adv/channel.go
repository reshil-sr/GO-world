package adv

import (
	"fmt"
)

func Chnl() {
	fmt.Println("Welcome to Channels")
	//step 1: make a channel of type string and pass to chCount
	ch := make(chan string)
	go func() {
		chCount("sheep", ch)
	}()
	msg := <-ch
	fmt.Println(msg)
	//this will print sheep one time and terminate

	//step 2: use a for loop instead
	/* for {
		msg, open := <-ch
		if !open { //if the channel is not open, ie when it is closed break out of the for loop
			break
		}
		fmt.Println(msg)
	} */
	//this will print sheep 5 times nd terminate with fatal error deadlock.
	//bcs the chCount finishes and the chnl function is still waiting to rcv on the channel
	// (will only send msg 5 times but here in for loop it waiting for the next msg),
	//Go identifies this in the runtime and throw deadlock
	//Inorder to fix this we use close(channel)
	//as a sender if we finished sending and we don't need the channel any more we can close the channel

	//step 3: More elegant way is to loop over the range of ch, so we don't need to manually check the channel is open or not

	for msg := range ch {
		fmt.Println(msg)
	}

}
func chCount(s string, ch chan string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(s)
		ch <- s //write string to channel s
	}
	close(ch)
}
