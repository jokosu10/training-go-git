package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello Feature A1")
	fmt.Println(printMessage("Feature B"))
}

func printMessage(name string) string {
	var message = "Hallo " + name
	return message
}
