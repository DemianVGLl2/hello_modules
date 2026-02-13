package hello_modules

import (
	"fmt"
	"math/rand"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hello %v from hello1", name)
	return message
}

func RandomHello() string {
	greetings := []string{
		"Hey, yo, %v",
		"Qué rollo con el pollo, %v",
		"What up, %v",
		"What up dog, my %v",
		"Holi crayoli, my %v",
	}

	return greetings[rand.Intn(len(greetings))]
}

func RandomGibberish() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	var gibberish string

	for i := 0; i < 8; i++ {
		index := rand.Intn(len(letters))
		gibberish[i] = letters[index]
	}

	return string(gibberish)
}
