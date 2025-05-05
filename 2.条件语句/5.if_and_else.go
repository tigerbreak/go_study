package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// 这种现象在所有带有副作用（如打印）的函数调用中都会出现。只要有函数在参数求值阶段有输出，都会先于主输出出现。
// Go 语言中，fmt.Println(a, b)会先计算所有参数的值，再把这些值一起输出。
// 也就是说，pow(3, 2, 10)和pow(3, 3, 20)会先后被调用，它们的返回值会被收集起来，然后fmt.Println才会把它们一起打印出来。
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
