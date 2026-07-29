package main

import ( 
	"fmt"
)

func superDiv(n int) bool{

	for i:=1;i<=20;i++{
		if n%i !=0 {
			return false
		}
	}

	return true
}

func main (){

	

	for num:=11;;num++ {

		if superDiv(num){
			fmt.Println(num)
			break
		}

	}
}
