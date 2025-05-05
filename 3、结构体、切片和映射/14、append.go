package main

import "fmt"

func main() {
	var a []int
	printSlice(a)

	//append 向切片添加元素
	a = append(a, 0)
	printSlice(a)

	//可以一次性添加多个元素
	a = append(a, 1, 2, 3)
	printSlice(a)
	
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
