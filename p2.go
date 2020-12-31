package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello to %s", name)
}

func main() {
	fmt.Printf("H")
	greet("ello")
}
