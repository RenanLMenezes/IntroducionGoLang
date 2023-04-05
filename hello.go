package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 3

func main() {

	showIntrodution()
	for {

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

		//use _ to ignore 1 return var
		// name, age := showNameAge()
		// fmt.Println(name, age)

		switch receiveCommand() {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing Logs...")
			printLog()
		case 0:
			fmt.Println("Bye bye...")
			os.Exit(0)
		default:
			fmt.Println("Choose a valid option")
		}
	}

}

// return 2 values
// func showNameAge() (string, int) {
// 	name := "Renan"
// 	age := 20
// 	return name, age
// }

func showIntrodution() {
	//variables
	// if u don't use a variable Go don't compile
	// u can declare without the type
	// u can u := instead var and without the type
	name := "Renan" //var name string
	//var age = 20      // if u don't initialize the variable the value is 0 ou null | var age int
	var version = 0.7 // var version float32 || float64

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

func startMonitoring() {
	fmt.Println("Monitoring...")
	// array
	// var sites [3]string
	// sites[0] = "https://github.com/RenanLMenezes"
	// slice
	//sites := []string{"https://github.com/RenanLMenezes", "https://supergeeks.com.br", "https://www.youtube.com"}
	//fmt.Println(len(sites), cap(sites)) //  show the length of the slide | capacity7
	sites := readFile()
	for i := 0; i < 3; i++ {
		//for i := 0; i < len(sites); i++
		for i, site := range sites {
			fmt.Println("Index", i, "Site:", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", nil)
	}

	if res.StatusCode == 200 {
		fmt.Println(site, "Success", res.StatusCode)
		writeLog(site, true)
	} else {
		fmt.Println(site, "Error", res.StatusCode)
		writeLog(site, false)
	}
}

func readFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro:", nil)
	}

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()

	return sites
}

func writeLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online :" + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLog() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	println(string(file))
}
