package main

import "fmt"

type fooBarer interface {
	foo() string
	bar() int
}

type myFooBar struct {
	fooBarer
}

func (myf *myFooBar) foo() string {
	return "foo string"
}

func showTheFooBar(fb fooBarer) {
	fmt.Println(fb.foo())
	//fmt.Println(fb.bar()) //this will panic since we only define foo method for myFooBar struct
}

func main() {
	fmt.Println("Welcome to Interfaces")
	s3Region := "aws-region"
	s3Endpoint := "s3Endpoint"
	path := "a/b/c/d.jpg"
	publishBucket := "public-bucket"
	txt := fmt.Sprintf(s3Endpoint, s3Region, fmt.Sprintf("%s/%s", publishBucket, path))
	fmt.Println(txt)
	//mock := myFooBar{}
	//fmt.Printf("%#v\n", mock)
	//showTheFooBar(&mock)
}
