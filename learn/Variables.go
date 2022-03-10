package main

// 1. All variables must be used else compile time error
// 2. lowerCase firt letter is package scope
// 3. Uppercase first letter is to export
// 4. there is no private scope
// 5. Capitalise acronym (HTTP)
// 6.  shorter name shorter live span
// 7. type conversion should be -> destinationType(variable) , no implicit conversion possible
// 8. use strConv package for strings

import (
	"fmt"
	"strconv"
)

var PIE float32 = 3.125
var (
	studentName string = "Nitish"
	studentCity string = "Gouda"
)

func main() {
	x := 3

	var j int = 42
	j = 20
	fmt.Println(x, j)

	for i := 0; i < 5; i++ {
		value := strconv.Itoa(i)
		fmt.Printf("%v, %T\n", i, i)
		fmt.Printf("%v, %T\n", value, value)
	}
}
