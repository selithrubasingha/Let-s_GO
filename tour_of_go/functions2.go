package main

import "fmt"

func add2(x, y int) int {
	return x + y
}

func sub2(x,y,z int) int {
	return x - y - z
}

func main4() {
	fmt.Println(add2(42, 13))
	fmt.Println(sub2(50,12,13))
}
