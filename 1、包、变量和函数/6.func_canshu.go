package main

import "fmt"

//x int, y int ,可以简写为 x,y int
func add(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(1, 2))
}
