package main

import (
	"fmt"
	"regexp"
	"sync"
)

// regxImagesgoodsFolder is used to check if a folder is imagesgoods folder.
var regxImagesgoodsFolder = regexp.MustCompile(`^([A-Z]{2})\/ST3\/[a-zA-Z]{2,}\/imagesgoods\/$`)

var count int

func main() {
	/* fmt.Println("hello Reshil")
	folderURI := "GU/ST3/SyncTest/imagesgoods/"
	dirDepthLessThan4 := len(strings.Split(folderURI, "/")) < 4
	isImagesgoods := IsImagesgoodsFolder(folderURI)
	fmt.Println(dirDepthLessThan4)
	fmt.Println(strings.Split(folderURI, "/"))
	fmt.Println(isImagesgoods) */

	cntrWithWg()
	brkVsReturn()

}

func IsImagesgoodsFolder(uri string) bool {
	return regxImagesgoodsFolder.MatchString(uri)
}

func cntrWithWg() {
	iter := 1000
	wg := new(sync.WaitGroup)
	wg.Add(iter)
	for i := 0; i < iter; i++ {
		go printCounter(wg)
	}
	wg.Wait()
	fmt.Println("the counter is:", count)
}
func printCounter(wg *sync.WaitGroup) {
	count++
	wg.Done()
}

func brkVsReturn() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("its 2")
			break // will break out of loop and print outside for loop
			//return // will break out of the loop and from the function
		}
		fmt.Println("still in for loop")
	}
	fmt.Println("outside for loop")
}
