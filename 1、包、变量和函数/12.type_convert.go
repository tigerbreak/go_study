package main

import (
    "fmt"    // 导入格式化输入输出包
    "math"   // 导入数学计算包
)

func main() {
    // 声明并初始化两个整型变量x和y
    var x, y int = 3, 4
    
    // 计算直角三角形的斜边长度
    // 1. x*x + y*y 计算两直角边的平方和
    // 2. float64() 将整型转换为浮点型，因为math.Sqrt需要浮点型参数
    // 3. math.Sqrt() 计算平方根
    var f float64 = math.Sqrt(float64(x*x + y*y))
    
    // 将浮点数f转换为无符号整型uint
    // 注意：这里会截断小数部分
    var z uint = uint(f)
    
    // 打印结果：x=3, y=4, z=5（勾股定理：3² + 4² = 5²）
    fmt.Println(x, y, z)
}