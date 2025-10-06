package main

import "fmt"

func main() {

	for a := 1; a <= 5; a++ {

		for b := 1; b <= a; b++ {
			fmt.Print(b, " ")
		}
		fmt.Println()
	}

    c:=1
	for {
		for b := 1; b <= c; b++ {
			fmt.Print(c, " ")
		}
		fmt.Println()
		c++
		if c>5{
			break
		}
	}
	
}
