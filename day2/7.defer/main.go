package main

import (
	"log"
	"time"
)

func main() {
	//n := [5]int{1, 2, 3, 4, 5}
	//for _, i := range n {
	//	defer fmt.Println(i)
	//}

	start := time.Now()
	log.Printf("开始时间为：%v", start)

	// defer的时候时间已经进去，所以打印出来不是3s
	//defer log.Printf("时间差：%v", time.Since(start))

	// 可以defer一个匿名函数
	defer func() {
		log.Printf("时间差：%v", time.Since(start))
	}()
	time.Sleep(3 * time.Second)

	log.Printf("函数结束")
}
