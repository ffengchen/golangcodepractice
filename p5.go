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

func seeker() {
	highest := 100
	lowest := 1
	for guess := 0; guess < 10; guess++ {
		mid := (highest + lowest) / 2
		res := ask(mid)
		if res == "Yes" {
			fmt.Println("Yes")
		}
		fmt.Println(res)
		if res == "Higher" {
			lowest = mid + 1
		} else {
			highest = mid - 1
		}

	}
}

func ask(answer int) string {

	fmt.Printf("Guess a number")
	for {
		g := askInt("Take a guess")
		if g == answer {

			return "Yes"
		}
		if g < answer {
			return "Higher"
		} else {
			return "Lower"
		}
	}
}

func main() {
	seeker()
}
