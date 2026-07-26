package main

import "fmt"

// -----------------------------------------
// 1. Define the Struct for our Maps
// -----------------------------------------
type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println("=== 1. Making Maps with 'make' ===")
	makingMaps()

	fmt.Println("\n=== 2. Map Literals (Explicit) ===")
	mapLiteralsExplicit()

	fmt.Println("\n=== 3. Map Literals (Implicit / Short) ===")
	mapLiteralsImplicit()

	fmt.Println("\n=== 4. Mutating Maps ===")
	mutatingMaps()
}

// -----------------------------------------
// 1. Initialization using 'make'
// -----------------------------------------
func makingMaps() {
	// A map's zero value is nil. You cannot add keys to a nil map!
	// You MUST use make() to initialize it before assigning values.
	var m map[string]Vertex
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println("Bell Labs Coordinates:", m["Bell Labs"])
}

// -----------------------------------------
// 2. Map Literals (Explicit typing)
// -----------------------------------------
func mapLiteralsExplicit() {
	// Map literals allow you to populate the map right when you declare it.
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println("Explicit Map:", m)
}

// -----------------------------------------
// 3. Map Literals (Implicit typing)
// -----------------------------------------
func mapLiteralsImplicit() {
	// Go is smart! If the map expects a 'Vertex' as the value, 
	// you can drop the word 'Vertex' inside the literal.
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println("Implicit Map (Cleaner):", m)
}

// -----------------------------------------
// 4. Mutating Maps (Insert, Delete, Check)
// -----------------------------------------
func mutatingMaps() {
	m := make(map[string]int)

	// 1. Insert a value
	m["Answer"] = 42
	fmt.Println("Inserted value:", m["Answer"])

	// 2. Update a value
	m["Answer"] = 48
	fmt.Println("Updated value:", m["Answer"])

	// 3. Delete a value
	delete(m, "Answer")
	// If a key doesn't exist, Go returns the "zero value" for that type (0 for int)
	fmt.Println("Value after delete:", m["Answer"]) 

	// 4. The "Comma Ok" idiom (Checking if a key exists)
	// 'v' gets the value (or 0 if missing). 'ok' gets a boolean (true if exists, false if missing)
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}