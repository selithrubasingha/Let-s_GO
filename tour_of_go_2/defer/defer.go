package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== 1. The Basic Defer ===")
	// This will get pushed to the end of the main() function
	defer fmt.Println("world (I was deferred!)")
	fmt.Println("hello")
	
	fmt.Println("\n=== 2. The Defer Stack (LIFO) ===")
	fmt.Println("Stacking the plates...")
	// We loop from 1 to 3, but defer the prints
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("Taking off plate %d\n", i)
	}
	fmt.Println("Done stacking! Now taking them off in reverse...")

	fmt.Println("\n=== 3. Real-World Simulation ===")
	// Calling another function to show how defer cleans things up
	doFileStuff()
	
	fmt.Println("\n=== Main function is ending NOW! ===")
	// Right after this line, all the deferred statements in main() will fire!
}

// A separate function to simulate opening and closing a file
func doFileStuff() {
	fmt.Println("[File] Opening 'passwords.txt'...")
	
	// We defer the close immediately so we don't forget it
	defer fmt.Println("[File] Closing 'passwords.txt'... (Safe and clean!)")
	
	fmt.Println("[File] Reading the file...")
	fmt.Println("[File] Doing some heavy hacker processing...")
	
	// The deferred file close will trigger exactly when this function finishes!
}