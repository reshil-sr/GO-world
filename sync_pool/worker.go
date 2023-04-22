package main

import (
	"fmt"
	"log"
	"net/http"
)

type Results struct {
	status int
}

type Site struct {
	url string
}

func main() {
	//crawler to crawl through a set of url and returns its status code
	fmt.Println("")
	webUrls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.youtube.com",
		"https://go.dev",
		"https://go.dev",
		"https://reshil.dev",
	}
	size := len(webUrls)
	jobs := make(chan Site, size)
	ResultChan := make(chan Results, size)

	for i := 0; i <= 3; i++ {
		go crawl(i, jobs, ResultChan)
	}

	go func() {
		for _, weburl := range webUrls {
			jobs <- Site{url: weburl}
		}
		close(jobs)
	}()
	for i := 1; i <= size; i++ {
		/* res, open := <-ResultChan
		fmt.Printf("%v\n", open)
		if !open {
			return
		} */
		res := <-ResultChan
		fmt.Printf("The response status is %d:\n", res.status)
	}
	/* for res := range ResultChan {
		fmt.Printf("The response status is %d:\n", res.status)
	} */
	/* close(ResultChan)
	<-ResultChan */
	fmt.Println("I am done")
}

func crawl(wId int, job <-chan Site, rslt chan<- Results) {
	fmt.Printf("The worker id :%d \n", wId)
	for site := range job {
		resp, err := http.Get(site.url)
		if err != nil {
			log.Println(err.Error())
			rslt <- Results{status: 404}
			continue
		}
		rslt <- Results{status: resp.StatusCode}
	}
}
