package main

import (
	"fmt"
)

func main(){

	for i:=1;i<=1000;i++{
		for j:=i;j<=1000;j++{
			for k:=j;k<=1000;k++{
				if i+j+k==1000 && i*i+j*j==k*k{
					fmt.Println(i*j*k)
				}
			}
		}
	}
}
