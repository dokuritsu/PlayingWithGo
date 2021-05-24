package main

import (
	"fmt"

	"rsc.io/quote"
)

// package for formatting text => printing to console

// implementing main function => NOTE TO SELF: PACKAGE MUST BE CLASSIFIED AS MAIN TO RUN AS MAIN FUNCTION
func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
}
