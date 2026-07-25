package main

import (
	"fmt"
	"math/rand"
)

func main3() {
	// -----------------------------------------
	// 1. Basic 'if'
	// -----------------------------------------
	score := 85
	if score >= 80 {
		fmt.Println("1. Basic if: Great job, you scored high!")
	}

	// -----------------------------------------
	// 2. Standard 'if else'
	// -----------------------------------------
	number := 7
	if number%2 == 0 {
		fmt.Println("2. if else: The number is even")
	} else {
		fmt.Println("2. if else: The number is odd")
	}

	// -----------------------------------------
	// 3. 'if' with a short statement
	// -----------------------------------------
	// Go lets you declare a variable right before the condition.
	// Variables declared this way only exist inside the if/else block!
	if n := rand.Intn(100); n > 50 {
		fmt.Printf("3. Short statement: %d is greater than 50\n", n)
	} else {
		fmt.Printf("3. Short statement: %d is 50 or less\n", n)
	}

	// If you tried to print 'n' here, Go would throw an error 
	// because 'n' stops existing the moment the if/else block ends.
}