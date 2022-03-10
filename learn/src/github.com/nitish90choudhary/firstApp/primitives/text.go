package main

import "fmt"

func main() {

	str := "this is a string "
	str2 := "this is another string"
	fmt.Printf("%v , %T\n", string(str[3]), str[3])
	fmt.Printf("%v \n", str+str2)
}
