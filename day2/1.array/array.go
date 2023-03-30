package main

import "fmt"

func main() {
	//a := [...]int{1, 2}
	//a[0] = 10
	//for _, val := range a {
	//	fmt.Println(val)
	//}

	// 多维数组
	arr := [2][2]int{{1, 2}, {2, 3}}
	for _, val := range arr {
		fmt.Println(val)
	}
}
