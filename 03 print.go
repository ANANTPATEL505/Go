package main
import "fmt"
func main(){
	name:="anant";
	age:=20
	height:=6.2

	fmt.Println("my name is",name,"and i am",age,"years old ")
	fmt.Printf("my name is %s and i am %d years old\n",name,age)

	fmt.Printf("%T\n",name)
	fmt.Printf("%T\n",age)
	fmt.Printf("%T\n",height)
}
