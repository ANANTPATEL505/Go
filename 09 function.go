package main
import (
	"fmt"
)

func f1(){
		fmt.Println("hello")
	}

func f2(a int){
		fmt.Println(a)
	}

func f3(a int ,b int){
	fmt.Println(a+b)
}

func f4(a int ,b int) int{
	return a+b
}	

func main() {
	
	f1()
	f2(2)
	f3(5,5)
	sum:=f4(10,10)
	fmt.Println(sum)
}
