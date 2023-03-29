package main

import "fmt"

func main() {
	//x := 1
	//if x == 0 {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("false")
	//}
	//fmt.Println("down")
	//if a := 10; a > 50 {
	//	fmt.Println("true")
	//}

	//// for
	//sum := 0
	//for i := 1; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//
	////while
	//n := 10
	//for n > 0 {
	//	n--
	//	fmt.Println(n)
	//}
	//
	//// return结束循环,不执行后面的
	//n = 10
	////for n > 0 {
	////	return
	////}
	////fmt.Println("down")
	//
	//// break会执行
	//for n > 0 {
	//	break
	//}
	//fmt.Println("down")

	// goto
	//	for i := 0; i < 10; i++ {
	//		for j := 0; j < i; j++ {
	//			if j == 2 {
	//				goto toHere
	//			}
	//		}
	//	}
	//
	//	return
	//toHere:
	//	fmt.Println("ok")

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Println()
	}

	m := map[string]int{
		"xyx": 1,
	}
	for key, val := range m {
		fmt.Println(val)
		m[key] = 2
	}
	fmt.Println(m["xyx"])
}
