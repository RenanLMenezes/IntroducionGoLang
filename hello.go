package main

import (
	"fmt"
	"os"
)

func main() {

	showIntrodution()

	showMenu()

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Showing Logs...")
	// } else if command == 0 {
	// 	fmt.Println("Bye bye...")
	// } else {
	// 	fmt.Println("Choose a valid option")
	// }

	//comando := receiveCommand()

	switch receiveCommand() {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Bye bye...")
		os.Exit(0)
	default:
		fmt.Println("Choose a valid option")
		os.Exit(-1)
	}
}

func showIntrodution() {
	//variables
	// if u don't use a variable Go don't compile
	// u can declare without the type
	// u can u := instead var and without the type
	name := "Renan" //var name string
	//var age = 20      // if u don't initialize the variable the value is 0 ou null | var age int
	var version = 0.2 // var version float32 || float64

	//print
	fmt.Println("Hello World! And", name)
	fmt.Println("This program is on the version", version)
}

func showMenu() {
	fmt.Println("\n1 - Start monitoring")
	fmt.Println("2 - View the logs")
	fmt.Println("0 - Exit")
}

func receiveCommand() int {
	fmt.Print("\nChoose your option: ")
	//func to receive input
	var command int
	fmt.Scan(&command) // or Scanf("%d", &commando)

	fmt.Println("The command is", command)
	//fmt.Println("The address of command is", &command) //var address

	return command
}
