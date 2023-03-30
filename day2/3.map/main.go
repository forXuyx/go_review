package main

import (
	"fmt"
	"sync"
)

func main() {
	//allMap := make(map[string]int)
	//allMap["beijing"] = 1
	//allMap["shanghai"] = 2
	//allMap["hangzhou"] = 3
	////delete(allMap, "shanghai")
	//allMap = make(map[string]int)
	//for k, v := range allMap {
	//	fmt.Println(k, v)
	//}

	// 以下代码线程不安全
	//m := make(map[int]int)

	//go func() {
	//	for {
	//		m[1] = 1
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		_ = m[1]
	//	}
	//}()
	//
	//for {
	//
	//}

	// 线程安全的map
	var m sync.Map
	go func() {
		for {
			m.Store(1, 1)
		}
	}()

	go func() {
		for {
			fmt.Println(m.Load(1))
		}
	}()

	for {

	}
}
