package main

import (
	"fmt"
	"log"

	"../greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// An array of names.
	names := []string{"Mark", "Mich", "Micoh"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
	for _, message := range messages {
		fmt.Println(message)
	}
}
