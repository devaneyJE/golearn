package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for generating and english greeting
	return "Hi, there!"
}

func (spanishBot) getGreeting() string { // note: can omit receiver sb if not in use; still need proper type
	// very custom logic
	return "Hola!"
}
