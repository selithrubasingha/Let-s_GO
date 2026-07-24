package main

import "fmt"

func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main8() {
	fmt.Println(split2(17))
}
