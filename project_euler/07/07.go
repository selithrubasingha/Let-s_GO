package main

import (
	"fmt"
	"math"
)


func isPrime(n int)bool {
	if n<2 {
		return false
	}

	limit :=  int(math.Ceil((math.Sqrt(float64(n)))))

	for i :=2;i<= limit;i++{

		if n%i==0 {
			return false
		}
	}

	return true
}

func main (){

	count :=1
	num := 1
	for {
		if isPrime(num){
			count++

			if count==10001 {
				fmt.Println(num)
				break
			}

		}

		num++

	}

}














