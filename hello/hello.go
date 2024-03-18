package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("woo")
	fmt.Println(message)
}
