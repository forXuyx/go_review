package main

//func main() {
//	//a, b := 1, 2
//	//c := tes(compareAB(a, b))
//	//fmt.Println(c)
//
//	a, b := 1, 2
//
//	swap(&a, &b)
//
//	fmt.Println(a, b)
//
//	s := []int{1, 2, 3}
//	res := test("sum:%d", s...)
//	fmt.Println(res)
//
//	// 匿名函数
//	getSqrt := func(a float64) float64 {
//		return math.Sqrt(a)
//	}
//	fmt.Println(getSqrt(4))
//	d := func(a float64) float64 {
//		return math.Sqrt(a)
//	}(4)
//	fmt.Println(d)
//}
//
//func test(s string, n ...int) string {
//	x := 0
//	for _, i := range n {
//		x += i
//	}
//	return fmt.Sprintf(s, x)
//}
//
//func swap(a, b *int) {
//	*a, *b = *b, *a
//}

//func testFu(compareAB func() bool, a, b int) bool {
//	return compareAB(a, b)
//}
//
//func compareAB(a, b int) bool {
//	return a > b
//}
//
//func visit(list []int, f func(int)) {
//	for _, v := range list {
//		f(v)
//	}
//}
//
//func main() {
//	visit([]int{1, 2, 3}, func(i int) {
//		fmt.Println(i)
//	})
//}

func playerGen(name string) func() (string, int) {
	hp := 150

	return func() (string, int) {
		return name, hp
	}
}

func main() {

}
