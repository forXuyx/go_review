package main

import (
	"fmt"
	"strings"
)

//func main() {
//	s1 := "localhost:8080"
//	fmt.Println(s1)
//
//	strByte := []byte(s1)
//
//	strByte[len(s1)-1] = '1'
//	s2 := string(strByte)
//	fmt.Println(s2)
//}

func main() {
	s1 := "hello,柠檬鱼的cpp之旅"
	oldStr := "cpp"
	newStr := "golang"
	index := strings.Index(s1, oldStr)
	oldLen := len(oldStr)
	start := s1[:index]
	end := s1[index+oldLen:]

	str := start + newStr + end
	fmt.Println(str)
}
