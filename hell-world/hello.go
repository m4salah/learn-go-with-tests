package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenshHelloPrefix = "Bonjour, "
const explinationMark = "!"

func greetingPrefix(lang string) string {
	switch lang {
	case "Frensh":
		return frenshHelloPrefix
	case "Spanish":
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
}
func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name + explinationMark
}

func main() {
	fmt.Println(Hello("world", ""))
}
