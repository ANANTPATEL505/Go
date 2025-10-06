package main

import (
	"fmt"
)

func main() {
	arr :=[]int{1,2,3,4,5}
	arr[1]=3

	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	slice:=[]int{10,20,30}
	slice1:=append(slice,40)
	fmt.Println(slice1)
}
