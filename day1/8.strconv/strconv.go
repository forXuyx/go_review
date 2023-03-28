package main

import (
	"fmt"
	"strconv"
)

func main() {
	newStr := "12"
	intVal, _ := strconv.Atoi(newStr)
	fmt.Printf("%d\n", intVal)

	str := "3.14159"
	floatVal, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("%T %f\n", floatVal, floatVal)
}
