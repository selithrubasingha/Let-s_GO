package main

import (
	"fmt"
	"math"

)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i :=2 ;i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 600851475143

	limit := int(math.Sqrt(float64(n)))

	sum := 0

	for i := 2; i<= limit; i++ {
		if n%i == 0 && isPrime(i) {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println(sum)


}
