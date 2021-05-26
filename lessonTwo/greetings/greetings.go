package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// Create a message with random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets inital values for variables used in the function
// used this to seed rand package with current time
// init functions are automatically executed at start up
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting msgs.
// Returned msg is random
func randomFormat() string {
	// slice of msg formats
	// []string indicates that size of the array is dynamic
	formats := []string{
		"Hi, %v. Welcome!",
		"Oh, it's you again, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
