package main
import "fmt"
func main(){
	var num int =10;
	var num2 =20;
	name:="hello";
	fmt.Println(num,num2,name);

	name,age:="anant",20;
	fmt.Println(name,age);

	name,age1:="patel",40;
	fmt.Println(name,age1);

	abc:="le bhai";
	_=abc; //silent

	var(
		school string;// empty
		class int; //defalut 0
		gender bool;//default false
	)
	fmt.Println(school,class,gender)

	var x,y,z int 
	fmt.Println(x,y,z)
}
