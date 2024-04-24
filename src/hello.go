package main

import "fmt"

const(
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	spanish = "Spanish"
	french = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := getPrefix(language)

	return prefix + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}