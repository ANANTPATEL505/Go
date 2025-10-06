package main

import (
	"fmt"
)

func main() {
	
	var str string ="hello my name is anant" + " hi"
	fmt.Println(str+" hi")

	fmt.Println(str[0]) //Ascii value h=104

	str="hello world"
	fmt.Println(str[6:11])
}
