package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// Return a greeting that embeds the name in a message.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a velue that embeds the name
	// in a greeing message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
