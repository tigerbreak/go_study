//场景一：函数参数传递
func addOne(n int) {
    n = n + 1
}

func main() {
    x := 10
    addOne(x)
    fmt.Println(x) // 输出10，x没有被修改
}
//用指针

func addOne(n *int) {
    *n = *n + 1
}

func main() {
    x := 10
    addOne(&x)
    fmt.Println(x) // 输出11，x被修改了
}

//高效处理大结构体数据
type BigStruct struct {
    Data [1000000]int //1000000个int
}

func process(bs *BigStruct) { //bs是BigStruct的指针
    bs.Data[0] = 100	//
}

func main() {
    var b BigStruct
    process(&b) //传入b的地址
    fmt.Println(b.Data[0]) // 输出100 
}