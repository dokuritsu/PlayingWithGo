package main

import (
	"fmt" // formatting text

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Alejandro")
	fmt.Println(message)
}
