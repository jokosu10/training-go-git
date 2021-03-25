package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello Feature A1")
	fmt.Println(printMessage("Feature B"))
	fmt.Println("Hello Feature A2")
	fmt.Println("Hello Feature B2")
}

// function to print name
func printMessage(name string) string {
	var message = "Hallo " + name
	return message
}
