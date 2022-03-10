package main

/*
integer = uint , int , int8, int16, int32 , int64

int +int8 -> leads to compile time error

complex64 = 2 +3i
and complex128
real() and imag() to pull real and imaginary parts
complex(5,6) make 5+6i
*/

import "fmt"

func main() {
	var isValid bool
	var num int
	isDay := false
	fmt.Printf("%v %T\n", isDay, isDay)

	if isDay {
		fmt.Println("day")
	} else {
		fmt.Println("night")
	}
	fmt.Println(isValid)
	fmt.Println(num)

	var x int = 8
	var y int8 = 4
	//fmt.Println("sum is ", x+y) //mismatched type compile time error
	fmt.Println("sum is ", x+int(y))

	var cnum1 complex64 = 2 + 3i
	var cnum2 complex64 = 2 - 0i
	fmt.Println("complex num sum is ", cnum1+cnum2)
}
