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

// Hellos returns a map that associates each of the named ppl
// with a greeting msg.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with msgs.
	messages := make(map[string]string)
	// Loop through received slice of names, calling
	// the Hello function to get a msg for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved msg with name
		messages[name] = message
	}
	return messages, nil
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