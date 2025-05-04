package main

import "fmt"

// var 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。 整型（int）：默认值是 0
//布尔型（bool）：默认值是 false
//字符串（string）：默认值是 ""（空字符串
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
