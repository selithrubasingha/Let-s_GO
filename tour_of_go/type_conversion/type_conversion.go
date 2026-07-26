package main

import (
	"fmt"
	"math"
)

func main12() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	
	var l,m int = 45 , 56
	
	var whoah float64 = math.Sqrt(float64(l*l+m*m))
	
	fmt.Println(whoah)
	
}
