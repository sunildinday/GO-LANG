package main

import (
	"fmt"

	"rsc.io/quote"
)

//"hello world"
func Hello() string {
	return quote.Hello()
}
func main() {
	fmt.Println("this is the entry point of the code")
}
