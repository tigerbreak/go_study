package main

import "fmt"
//:= 语法是短变量声明，用于声明和初始化局部变量。它只能在函数内部使用。 var 语法用于声明全局变量。
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
