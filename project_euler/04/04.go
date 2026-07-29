package main

import (
	"fmt"
	"strconv"
)


func isPalindrome(n int) bool {
	str := strconv.Itoa(n)

	for i,j:=0,len(str)-1;i<j;i,j=i+1,j-1{
		if str[i]!=str[j]{
			return false

		}
	}

	return true
}

func main(){

	var breakloop bool = false
	for i:=1000;i>0;i--{
		for j:=1000;j>0;j--{
			var num = i *j

			if  isPalindrome(num) {
				fmt.Println(num)
				breakloop = true
				break
			}


		}

		if breakloop {
			break
		}
	}

}
