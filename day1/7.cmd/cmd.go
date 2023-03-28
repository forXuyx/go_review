package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "运行模式可以设置为fast")

func main() {
	// 解析命令行
	flag.Parse()
	fmt.Printf(*mode)
}
