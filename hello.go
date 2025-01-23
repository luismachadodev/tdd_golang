package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name, idiome string) string {
	if name == "" {
		name = "World"
	}

	if idiome == "espanhol" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
