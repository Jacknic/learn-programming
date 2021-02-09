package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := rand.Intn(10), rand.Intn(10)
	fmt.Printf("a=%d,b=%d\n", a, b)
	if a > b {
		fmt.Println("a 大于 b")
	} else {
		fmt.Println("a 不大于 b")
	}
	if a < 9 {
		if a > 5 {
			fmt.Println("a 比五大且比8小")
		}
		fmt.Println("样式1")
		switch {
		case a == 1:
			fmt.Println("中文数字：", "壹")
		case a == 2:
			fmt.Println("罗马数字：", "II")
		default:
			fmt.Println("阿拉伯数字：", a)
		}
		fmt.Println("样式2")
		switch a {
		case 1:
			fmt.Println("中文数字：", "壹")
		case 2, 3:
			fmt.Println("罗马数字：", "II或III")
		default:
			fmt.Println("阿拉伯数字：", a)
			break
		}
	}
	for i := 0; ; i++ {
		if i > 10 {
			fmt.Println("\n跳出循环")
			break
		}
		if i == 4 {
			continue
		}
		fmt.Print(i)
	}
	fmt.Println()
}
