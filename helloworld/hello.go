package main

import "fmt"

const (
	russian = "Russian"
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	russianHelloPrefix = "Привет, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

//private func (starts with lowercase letter) with a NAMED return value
func greetingPrefix(language string) (prefix string) {
	switch language {
	case russian:
		prefix = russianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}
