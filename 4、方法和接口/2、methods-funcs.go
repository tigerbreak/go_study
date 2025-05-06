package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
//函数：func 函数名(参数列表) 返回值类型
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
