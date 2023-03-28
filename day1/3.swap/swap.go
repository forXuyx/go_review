package main

import "fmt"

func main() {
	//a, b := 1, 2
	//var temp int
	//fmt.Println(a, b)
	//temp = b
	//b = a
	//a = temp
	//fmt.Println(a, b)

	a, b := 1, 2
	b, a = a, b
	fmt.Println(a, b)
}
