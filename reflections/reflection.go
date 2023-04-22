package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Welcome to GO Reflection!!")

	type details struct {
		name string `env:"name,default='res',required=true"`
		age  int
	}

	type myType string

	person := details{
		name: "Reshil",
		age:  34,
	}
	//var str myType = "I am a newbie in go" //can also declare like this or as below
	str := myType("Hello Go!")

	showTypeAndKindDetails(person)
	showTypeAndKindDetails(str)
	showFieldDetails(str)
	showFieldDetails(person)
}

func showTypeAndKindDetails(intf interface{}) {
	ty := reflect.TypeOf(intf)
	kin := ty.Kind()
	fmt.Printf("Type %v is of kind:%v\n", ty, kin)
}

func showFieldDetails(intf interface{}) {
	ps := reflect.ValueOf(intf)
	//if pass person : //Type of ps is reflect.Value and its value is:{Reshil 34}
	//if pass details{} : //Type of ps is reflect.Value and its value is:{"" 0}
	fmt.Printf("Type of ps is %T and its value is:%v\n", ps, ps)
	fmt.Println(ps)        //{Reshil 34}
	fmt.Println(ps.Type()) //main.details
	fmt.Println(ps.Kind()) //struct
	if ps.Kind() == reflect.Struct {
		fmt.Println(ps.NumField()) //2
		for i := 0; i < ps.NumField(); i++ {
			fmt.Println(ps.Field(i))
		}
	}
}
