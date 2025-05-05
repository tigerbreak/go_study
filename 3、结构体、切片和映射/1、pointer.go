package main
import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //& 操作符会生成一个指向其操作数的指针。 &i 表示i的内存地址

	fmt.Println(*p) // 通过指针读取i的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}