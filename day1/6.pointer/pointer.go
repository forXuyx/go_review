package main

import "fmt"

// 指针有两个核心概念
// 类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，无需拷贝数据，类型指针不能进行偏移和运算
// 切片，由指向起始元素的原始指针、元素数量和容量构成

func add(a int) {
	a = a + 1
}

func add1(a *int) {
	*a = 1000
}

func main() {
	//a := 100
	//add(a)
	//fmt.Println(a)
	//add1(&a)
	//fmt.Println(a)

	// new函数创建指针
	ptr := new(string)
	*ptr = "柠檬鱼"
	fmt.Printf("%s", *ptr)
}
