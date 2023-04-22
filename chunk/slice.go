package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("chunk slice")
	urls := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "9", "10", "11", "12", "13"}
	perLimit := 3
	for _, chunkURLs := range chunkSlice(urls, perLimit) {
		fmt.Println(len(chunkURLs))
		go chunkBySize(chunkURLs, perLimit)

	}
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

func chunkSlice(slice []string, sizePerChunk int) [][]string {
	var chunks [][]string
	for index := 0; index < len(slice); index += sizePerChunk {
		chunks = append(chunks, slice[index:min(index+sizePerChunk, len(slice))])
	}
	return chunks
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func chunkBySize(urls []string, limit int) {
	fmt.Println(urls)
}
