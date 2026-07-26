package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== 1. Arrays vs Slice Literals ===")
	sliceLiterals()

	fmt.Println("\n=== 2. Slices are References to Arrays ===")
	slicesAreReferences()

	fmt.Println("\n=== 3. Slice Bounds, Length (len), and Capacity (cap) ===")
	lengthAndCapacity()

	fmt.Println("\n=== 4. Nil Slices & Creating with 'make' ===")
	makeAndNilSlices()

	fmt.Println("\n=== 5. Appending to Slices ===")
	appendingSlices()

	fmt.Println("\n=== 6. Range Iteration ===")
	rangeIteration()

	fmt.Println("\n=== 7. 2D Slices (Tic-Tac-Toe) ===")
	twoDimensionalSlices()
}

// -----------------------------------------
// 1. Literals and Basics
// -----------------------------------------
func sliceLiterals() {
	// An Array has a fixed size [6]
	primesArray := [6]int{2, 3, 5, 7, 11, 13}
	// A Slice has no size in the brackets []
	s1 := primesArray[1:4] 
	fmt.Println("Sliced from array:", s1)

	// A Slice Literal builds the underlying array for you automatically
	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, true, false, true}
	
	// A slice of anonymous structs!
	s := []struct {
		i int
		b bool
	}{
		{2, true}, {3, false}, {5, true},
	}
	fmt.Println("Int slice:", q)
	fmt.Println("Bool slice:", r)
	fmt.Println("Struct slice:", s)
}

// -----------------------------------------
// 2. Modifying Slices Modifies the Array
// -----------------------------------------
func slicesAreReferences() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println("Original Array:", names)

	a := names[0:2] // [John Paul]
	b := names[1:3] // [Paul George]

	// Changing 'b' also changes 'names' and 'a' because they 
	// all look at the exact same underlying memory!
	b[0] = "XXX"
	fmt.Println("Slice a:", a)
	fmt.Println("Slice b:", b)
	fmt.Println("Mutated Array:", names)
}

// -----------------------------------------
// 3. Length, Capacity, and Reslicing
// -----------------------------------------
func lengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("Initial", s)

	// Slice to give it zero length (but capacity remains!)
	s = s[:0]
	printSlice("Zeroed length", s)

	// Extend its length back out
	s = s[:4]
	printSlice("Extended length", s)

	// Drop its first two values (This reduces capacity!)
	s = s[2:]
	printSlice("Dropped first two", s)
}

// -----------------------------------------
// 4. Nil Slices and 'make'
// -----------------------------------------
func makeAndNilSlices() {
	var nilSlice []int
	if nilSlice == nil {
		fmt.Printf("nilSlice is nil! len=%d cap=%d\n", len(nilSlice), cap(nilSlice))
	}

	// make() is how you create dynamically sized slices
	a := make([]int, 5)    // len=5, cap=5
	b := make([]int, 0, 5) // len=0, cap=5
	printSlice("make(a)", a)
	printSlice("make(b)", b)
}

// -----------------------------------------
// 5. Append (Dynamic Growth)
// -----------------------------------------
func appendingSlices() {
	var s []int // Starts as nil
	
	// append works perfectly on nil slices
	s = append(s, 0)
	printSlice("Appended 0", s)

	s = append(s, 1)
	printSlice("Appended 1", s)

	// You can append multiple items at once
	s = append(s, 2, 3, 4)
	printSlice("Appended 2,3,4", s)
}

// -----------------------------------------
// 6. Range Iteration
// -----------------------------------------
func rangeIteration() {
	pow := make([]int, 5)
	
	// We can use range just for the index (omitting the value)
	for i := range pow {
		pow[i] = 1 << uint(i) // bit shift: 2**i
	}

	// We can use '_' to ignore the index and just get the value
	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

// -----------------------------------------
// 7. 2D Slices
// -----------------------------------------
func twoDimensionalSlices() {
	// Slices can contain other slices (like an array of arrays)
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// --- Helper Function ---
func printSlice(label string, s []int) {
	fmt.Printf("%-18s -> len=%d cap=%d %v\n", label, len(s), cap(s), s)
}