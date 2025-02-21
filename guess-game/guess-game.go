package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		var guess int
		fmt.Println("Enter a number: ")
		guess = 10
		fmt.Scan(&guess)
		if guess != 0 {
			min, max := 1, 100
			randomNumber := rand.Intn(max-min) + min
			if randomNumber == guess {
				fmt.Printf("your guess (%v) is right", guess)
			} else {
				fmt.Println("corret answer: ", randomNumber)
			}
		} else {
			break
		}
	}

}
