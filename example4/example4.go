package main

import (
	"fmt"
)
//Counter function which returns a function closure which increments 
//the number variable when invoked
func counter() func() int{

	var number int

	return func() int {
		number++
		return number
	}
}

func main() {
	
	//firstCounter will increment the same number variable each time
	//it is invoked
	firstCounter := counter()
	fmt.Println("First Counter Values")
	fmt.Println(firstCounter())
	//1 is printed
	fmt.Println(firstCounter())
	//2 is printed

	secondCounter := counter()
	fmt.Println("Second Counter Values")
	fmt.Println(secondCounter())
	//1 is printed
	fmt.Println(secondCounter())
	//2 is printed

}