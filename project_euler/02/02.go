package main

import "fmt"

func main (){

	first := 1
	second := 2

	res:=2

    fib :=first+second

	for fib<4000000 {
		if fib%2==0 {
			res+= fib
		}

		first = second
		second = fib

		fib = first + second

	}





	fmt.Println(res)
}
