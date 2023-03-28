package main

import "fmt"

// var 变量名 变量类型
//var age int // 声明一个int变量

// 基本变量类型：bool, string, int, uint(无符号), byte(uint8的别名), rune(int32的别名), float, complex
// 所有内存在Go中均被初始化
//func main() {
//	fmt.Printf("%d\n", age)
//}

// 不指名类型
//var level = 1.0
//
//func main() {
//	fmt.Printf("%T\n", level)
//}

// 声明多个变量
//var (
//	a int
//	b float32
//	c string
//)
//
//func main() {
//	fmt.Printf("a = %d, b = %f, c = %s\n", a, b, c)
//}

// 简短格式(只能定义在内部)

//func main() {
//	i := "start"
//	fmt.Printf("%s\n", i)
//}

// 多重赋值
//func main() {
//	conn, err := net.Dial("tcp", "127.0.0.1:8000")
//	conn1, err := net.Dial("tcp", "127.0.0.1:8000")
//	fmt.Println(conn, err, conn1)
//}

// 匿名变量
//func main()  {
//	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
//	fmt.Println(conn)
//}

// 变量作用域
// 函数体以内为局部变量 以外为全局变量
// 局部变量和全部变量名称可以相同，但是局部变量会优先考虑
// 形参只有函数调用时才会生效，调用后会被销毁，未被调用时没有空间也没有实际值

//func main() {
//	a, b := 1, 2
//	c := sum(a, b)
//	fmt.Println(c)
//}
//
//func sum(a, b int) int {
//	return a + b
//}

// 浮点型表示 通常会优先使用int64
//func main() {
//	floatNum := 3.2
//	fmt.Printf("%.2f", floatNum)
//}

// 字符型byte
//func main() {
//	var a rune = '1'
//	var b rune = 'b'
//	var c rune = ' '
//	fmt.Println(unicode.IsDigit(a))
//	fmt.Println(unicode.IsLetter(b))
//	fmt.Println(unicode.IsSpace(c))
//}

// 使用飘号可以原样输出
//func main() {
//	var str = `hellp \n 666`
//	fmt.Println(str)
//}

//func main() {
//	var str1 string = "hello"
//	var str2 [5]byte = [5]byte{104, 101, 108, 108, 111}
//	fmt.Printf("%c, %c\n", str1[0], str2[1])
//}

// 传统拼接 直接加 效率会低一点 可以用bytes.Buffer类型 再使用writestring方法
//func main() {
//	var strBuilder bytes.Buffer
//	s1 := "hello"
//	s2 := "world"
//	strBuilder.WriteString(s1)
//	strBuilder.WriteString(s2)
//	fmt.Printf("%s", strBuilder.String())
//}

// 取字符串里面的中文
//func main() {
//	s := "hello 世界"
//	fmt.Printf("%s\n", string([]rune(s)[6]))
//	for _, str := range s {
//		fmt.Printf("%c\n", str)
//	}
//}

// print 结果写到标准输出 sprint以字符串形式返回

// 变量周期
// 全局变量：和整个程序的运行周期一致
// 局部变量：动态的，从创建开始，到不再被引用为止
// 形参和函数返回值：调用时被创建，函数结束后销毁

// go使用两种数据结构存放变量
// 1.堆：用于存放进程执行过程中动态分配的内存段
// 2.栈：用来存放局部变量

// 类型别名
type Myint int
type Myint2 = int

func main() {
	// 类型转换 有时候会有精度损失
	//a := 5.01
	//b := int(a)
	//fmt.Println(a, b)
	var a Myint = 2
	var b Myint2 = 2
	fmt.Printf("%T, %T\n", a, b)
}
