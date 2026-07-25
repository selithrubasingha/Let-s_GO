package main 

import "fmt"

func main(){

	sum := 0
	for i:= 0; i<10; i++ {
		sum +=i
	}

	fmt.Println(sum)

	sum2 := 0

	// for loop as a while loop
	for sum2 < 1000 {
		sum2 += 23
	}

	fmt.Println(sum2)

	// never ending loop
	for {

	}
}










