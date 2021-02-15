package main

import (
	"fmt"
)

//Function literal (anonymous function)
func main() {
	
	func() {
		fmt.Println(5)
	}()

}