//用一个学生来展示结构体
type Student struct {
    Name  string
    Age   int
    Score float64
}

func main() {
    s := Student{Name: "小明", Age: 18, Score: 95.5}
    fmt.Println(s)
    fmt.Println("姓名：", s.Name)
    fmt.Println("年龄：", s.Age)
    fmt.Println("成绩：", s.Score)
}
//用一个点来展示结构体
type Point struct {
    X int
    Y int
}

func main() {
    p1 := Point{X: 10, Y: 20}
    p2 := Point{X: 30, Y: 40}
    fmt.Println("第一个点：", p1)
    fmt.Println("第二个点：", p2)
}
//用一个订单系统来展示结构体
type Order struct {
    OrderID   string
    ItemName  string
    Quantity  int
}

func main() {
    order := Order{OrderID: "20240505", ItemName: "手机", Quantity: 2}
    fmt.Println(order)
}