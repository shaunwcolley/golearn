package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for a person
func Hello(name string) (string, error) {
	// return error if no name
	if name == "" {
		return "", errors.New("empty name")
	}


	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in function

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return randomly slected message format using random index
	return formats[rand.Intn(len(formats))]

}