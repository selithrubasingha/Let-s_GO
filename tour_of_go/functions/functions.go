package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x int , y int )int {
	return x - y

}

func multiply(x int , y int )int{
	return x * y	
}

func main3() {
	fmt.Println(add(42, 13))
	fmt.Printf("Subtractio ! : %v \n", subtract(42,13))
	fmt.Println(multiply(10,34))
}
