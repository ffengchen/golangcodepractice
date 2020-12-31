package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askInt(message string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		s := scanner.Text()
		n, err := strconv.Atoi(s)
		if err == nil {
			return n
		}
		fmt.Printf("Error:'%s' is not a number", s)
	}

}

func main() {
	answer := 35
	fmt.Printf("Guess a number")
	for {
		g := askInt("Take a guess")
		if g == answer {
			fmt.Printf("You Got it")
			return
		}
		if g < answer {
			fmt.Printf("Nope: Higher")
		} else {
			fmt.Printf("nope: Lower")
		}
	}

}
