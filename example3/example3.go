package main

import (
	"fmt"
)

//Function literal used as a closure in numberPrinter()
func numberPrinter(number int) {

	func(){
		//This function literal has access to the number variable,
		//which is defined outside of the function body
		fmt.Println(number)
	}()
}
func main() {

	numberPrinter(7)

}