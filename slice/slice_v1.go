package main

//ref https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b

import "fmt"

/*
When we pass a slice to a function as an argument the values of the slice are passed by reference (since we pass a copy of the pointer),
but all the metadata describing the slice itself are just copies.
type sliceHeader struct {
    Length   int
	Capacity int
    ZerothElement *byte
}

slice := sliceHeader{
    Length:   50,
	Capacity: 50,
    ZerothElement: &buffer[100],
}
*/
func main() {
	fmt.Println("Hello! Welcome to  Slice!")
	explainOne()
	explainTwo()
	explainThree()
	explainFour()
}

func explainOne() {
	fmt.Println("Expalin One")
	slice := []string{"a", "a"}
	func(slice []string) {
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}(slice)
	fmt.Println(slice)
}

func explainTwo() {
	/*
		We can modify the data of the slice in the literal function,
		however if the pointer to the data changes for any reason or the slice metadata is modified,
		the change can be partially or no visible at all to the outside function.
	*/
	fmt.Println("Expalin Two")
	slice := []string{"a", "a"} //len: 2, cap: 2

	func(slice []string) {
		/*
			append:
			if there is enough capacity, the underlying array is reused;
			if not, a new underlying array is allocated and the data is copied over.
		*/
		/*
			no capacity(2), so a new underlying array is allocated,hence the pointer to the data changes and change wont visible to the outside function
			the pointer in the slice copy got overwritten
		*/
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice) //[b b a]
	}(slice)
	fmt.Println(slice) //[a,a]
}
func explainThree() {
	/*
		Moving the append after the manipulation of the slice,
	*/
	fmt.Println("Expalin Three")
	slice := []string{"a", "a"} //len: 2, cap: 2

	func(slice []string) {
		/*
			append:
			if there is enough capacity, the underlying array is reused;
			if not, a new underlying array is allocated and the data is copied over.
		*/
		slice[0] = "b"
		slice[1] = "b"
		/*
			no capacity(2), so a new underlying array is allocated,hence the pointer to the data changes and change wont visible to the outside function
			the pointer in the slice copy got overwritten
		*/
		slice = append(slice, "a")
		fmt.Println(slice) //[b b a]
	}(slice)
	/*
		Moving the append after the manipulation of the slice,
		we can notice that the behavior is different since the slice got reallocated after the manipulation of the values
		and the pointer is still pointing to the initial memory address.
	*/
	fmt.Println(slice) //[b,b]
}

func explainFour() {
	fmt.Println("Expalin Four")
	slice := make([]string, 2, 3) //len 2, cap 3
	func(slice []string) {
		slice = append(slice, "a") //cap for one more string, hence array is not allocated again and the pointer stays the same.
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice) //[b b a]
	}(slice)
	fmt.Println(slice)
	/*[b b] since the initialized and passed slice is of length 2, as mentioned above
	all the metadata describing the slice itself are just copies.
	*/
}
