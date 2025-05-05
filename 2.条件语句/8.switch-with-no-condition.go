package main

import (
	"fmt"
	"time"
)

// 这种形式能将一长串 if-then-else 写得更加清晰。
func main() {
	t := time.Now()
	fmt.Println("t的值为", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}
