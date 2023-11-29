package pointers

// invalid operation: cannot indirect res (variable of type []*Test)
// indicates that you are attempting to dereference a variable of type []*Test incorrectly.

import "fmt"

type Test struct {
	Name string
}

func main() {
	var res []*Test

	res = append(res, &Test{Name: "John"})
	res = append(res, &Test{Name: "Doe"})

	fmt.Println(*res) // Error: invalid operation: cannot indirect res (variable of type []*Test)
}

/*
In Go, the * operator is used to dereference a pointer, allowing you to access the value that the pointer points to.
However, in this case, you are trying to dereference a slice of pointers, which is not a valid operation.
In this example, we have a slice res of type []*Test (a slice of pointers to Test objects).
When you try to dereference the res variable using the * operator, as in *res, you get the mentioned error.

To access the elements of the res slice, you need to index it using square brackets []. Here's an updated version of the code:
*/

func workingWay() {
	fmt.Println(res[0].Name) // John
	fmt.Println(res[1].Name) // Doe
}
