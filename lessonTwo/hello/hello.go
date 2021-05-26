package main

import (
	"fmt" // formatting text
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Dokuritsu")

	if err != nil {
		log.Fatal(err) //Fatal = Print equivalent
	}

	fmt.Println(message)
}
