package main

import "fmt"

// -----------------------------------------
// 1. Defining Structs
// -----------------------------------------
type Vertex struct {
	X, Y int // You can group fields of the same type on one line!
}

type Man struct {
	Name   string
	Age    int
	Number string
}

// -----------------------------------------
// Package-Level Variables & Struct Literals
// -----------------------------------------
var (
	v1 = Vertex{1, 2}  // Positional setup (type: Vertex)
	v2 = Vertex{X: 1}  // Named setup (Y is implicitly 0)
	v3 = Vertex{}      // Empty setup (X:0 and Y:0)
	p  = &Vertex{1, 2} // Direct pointer initialization (type: *Vertex)
)

func main() {
	fmt.Println("=== 1. Basic Initialization ===")
	// Positional initialization using the exact data you provided
	selith := Man{"Selith", 22, "07132323232"}
	fmt.Println("Man struct:", selith)
	
	// Printing the package-level variables
	fmt.Println("Vertex literals:", v1, p, v2, v3)

	fmt.Println("\n=== 2. Accessing and Modifying Fields ===")
	v := Vertex{1, 2}
	fmt.Println("Original Vertex v:", v)
	
	// Modifying a field using simple dot notation
	v.X = 4
	fmt.Println("v.X after modification:", v.X)

	fmt.Println("\n=== 3. Struct Pointers & Implicit Dereferencing ===")
	p2 := &v // p2 is now a pointer to v

	// THIS IS GO MAGIC:
	// In C/C++ you would need to write (*p2).X = 1e9 or p2->X = 1e9
	// Go is smart enough to implicitly dereference struct pointers for you!
	p2.X = 1e9 
	
	fmt.Println("Vertex v after pointer modification:", v)
}