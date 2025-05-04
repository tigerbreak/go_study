package main

import (
	"fmt"
	"math"
)

// 在引入一个包时，只能引用已经导出的名字。以大写开头，未导出的名字在该包外无法访问，如下面的Pi
func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.E)
}
