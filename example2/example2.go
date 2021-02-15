package main

import (
	"fmt"
)

//Function Literal assigned to a variable and invoked 
func main() {

	number := 6

	numberPrinter := func() {
		fmt.Println(number)
	}

	numberPrinter() 

}