package main

import (
	"fmt"
	"reflect"
)

func main() {
	//variables
	// if u don't use a variable Go don't compile
	// u can declare without the type
	// u can u := instead var and without the type
	name := "Renan"   //var name string
	var age = 20      // if u don't initialize the variable the value is 0 ou null | var age int
	var version = 0.2 // var version float32 || float64

	//print
	fmt.Println("Hello World! And", name, "your age is", age)
	fmt.Println("This program is on the version", version)
	fmt.Println("Type of variable version is", reflect.TypeOf(version)) // show the type of the variable
}
