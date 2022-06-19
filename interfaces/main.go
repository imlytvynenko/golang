package main

import "fmt"

// Here I'm going to show the simplest problem, why we need interfaces
// This example is contrived

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	printGreeting(eb)
}

// Go does not support function overloading
// Also, printGreeting in both cases practically identicall,
// And generally it does not matter which bot is going to passed into func
// There are a lot of different languages 
// and it's not good write new print func for each bot
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hi there!"
}
