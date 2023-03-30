package main

// 切片是对数组一个连续片段的引用
// slice [开始位置:结束位置]

func main() {
	//a := [4]int{1, 2, 3, 4}
	//fmt.Println(a[1:1])
	//var highBuilder [30]int
	//
	//for i := 0; i < 30; i++ {
	//	highBuilder[i] = i + 1
	//}
	//
	//fmt.Println(highBuilder[:20])
	//fmt.Println(highBuilder[3:12])
	//fmt.Println(highBuilder[10:])

	// 直接生成切片
	//var strList []string

	//var intList []int
	//
	//var numList = []int{}
	//
	//fmt.Println(len(strList), len(intList), len(numList))

	//strList = append(strList, "xyx")
	//fmt.Println(strList)

	// 动态创建切片
	//a := make([]int, 2)
	//b := make([]int, 2, 10)
	//fmt.Println(len(a), cap(a), len(b), cap(b))
	//a = append(a, 2)
	//fmt.Println(len(a), cap(a), len(b), cap(b))
	//fmt.Println(a)

	// copy函数：原切片的变化不会影响复制后切片的变化
	// 引用切片：原切片的更改会影响现有切片
}
