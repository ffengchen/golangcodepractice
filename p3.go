package main

import "fmt"

func greet() (string, int) {
	fmt.Printf("Hello, what's your name?")
	var uName string
	fmt.Scanf("%s", &uName)
	var age int
	fmt.Printf("How old are you?")
	fmt.Scanf("%d", &age)
	return uName, age
}

func sayGoodbye(name string) {
	name = name + " Noggin"
	fmt.Printf("Goodbye to %s", name)
}

func main() {
	name, age := greet()
	fmt.Printf("Name is %s, age is %d", name, age)
	sayGoodbye(name)
}
