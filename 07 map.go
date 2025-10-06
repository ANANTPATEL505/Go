package main

import (
	"fmt"
)

func main() {
	
	stud:=map[string]int{}
	stud["anant"]=20
	fmt.Println(stud)

	student:=map[int]string{
		1:"anant",
		2:"hello",
		3:"world",
	}
	student[7]="katil"
	fmt.Println(student)

	for key,value:=range student{
		fmt.Println(key,value)
	}

	address:=map[int]map[string]string{
		1:{
			"name":"anant",
			"city":"surat",
		},
		2:{
			"name":"hello",
			"city":"space",
		},
	}
	fmt.Println(address[1]["name"],address[1]["city"])
}
