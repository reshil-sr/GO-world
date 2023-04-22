package adv

import (
	"fmt"
)

func Conc() {
	fmt.Println("Welcome to concurrency world")
	//step 1: this will print sheep infinetly and never reach code to call count("fish")
	/* count("sheep")
	count("fish") */

	//step 2: Now we make the count(sheep) a go routine, so that will run in the background
	// as a result both prints
	/* go count("sheep")
	count("fish") */

	// step 3: Now we make both go rountines, so that both will run in the background and the expected behaviour is printing both
	// but the program exits without printing anything, this is bcs when the main go routine finishes the program exits
	// without considering the background running goroutines, here the background goroutined didn't get any time to execute

	// step 4: we add a 2 seconds delay, now it will run for 2 seconds and exit

	/* go count("sheep")
	go count("fish")
	time.Sleep(2 * time.Second) */

	//Step 5: wait group
	WtGrp()
}

/* func count(s string) {
	for {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 500)
	}
} */
