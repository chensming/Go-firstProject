package main

import "fmt"

func test() {
test:
	fmt.Println("test")
}

//这是一个错误演示，意思是想表达goto关键字不能跨函数跳转，也不能转至内层代码块中
func main() {
	for i := 0; i < 3; i++ {
	next:
		fmt.Println("i = ", i)
	}
	goto test
	goto next
}
