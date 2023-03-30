package main

import "fmt"

type Student struct {
	name string
	age  int
	sex  string
}

func (s *Student) GetName() {
	fmt.Printf("My name is %s\n", s.name)
	s.name = "1"
}

func (s *Student) Rename(newName string) {
	s.name = newName
}

func main() {
	// 一个类型加上他的方法等价一个类
	s := &Student{
		name: "xyx",
		age:  25,
		sex:  "boy",
	}
	s.GetName()
	s.Rename("yjn")
	s.GetName()
	s.GetName()
}
