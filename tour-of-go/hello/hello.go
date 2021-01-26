package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of predefineed Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and a line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to console and
	// exit the program

	if err != nil {
		log.Fatal(err)
	}

	// If no err, print the returned message to console
	fmt.Println(message)
}