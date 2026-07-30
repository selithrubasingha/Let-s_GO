package main

import (
	"fmt"
)


func main() {
	sum1 := 0
	sum2 := 0

	for i:=0;i<=100;i++{
		sum1 += i*i
		sum2 += i

	}

	final := sum2*sum2 - (sum1)

	fmt.Println(final)
}