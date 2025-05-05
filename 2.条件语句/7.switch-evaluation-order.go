package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	fmt.Println("today =", today)
	fmt.Println("today+0 =", today+0)
	fmt.Println("today+1 =", today+1)
	fmt.Println("today.Monday =", time.Monday)
	switch time.Monday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}
}
