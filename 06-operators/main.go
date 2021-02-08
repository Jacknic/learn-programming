package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//算术运算符
	arithmetic()
	//关系运算符
	relation()
	//逻辑运算符
	logic()
	//位运算符
	bit()
	//赋值运算符
	assignment()
	//其他运算符
	others()
}

func arithmetic() {
	fmt.Println("算术运算符------------------------")
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}

func relation() {
	fmt.Println("关系运算符------------------------")
	rand.Seed(time.Now().UnixNano())
	a, b := rand.Intn(10), rand.Intn(10)
	fmt.Printf("a=%d,b=%d\n", a, b)
	if a > b {
		fmt.Println("a 大于 b")
	}
	if a >= b {
		fmt.Println("a 大等于 b")
	}
	if a < b {
		fmt.Println("a 小于 b")
	}
	if a <= b {
		fmt.Println("a 小等于 b")
	}
	if a == b {
		fmt.Println("a 等于 b")
	}
}

func logic() {
	fmt.Println("逻辑运算符------------------------")
	randBool := func() bool { return rand.Int()&1 == 0 }
	finished, end := randBool(), randBool()
	fmt.Printf("finished=%v,end=%v\n", finished, end)
	if finished && end {
		fmt.Println("完成并结束√")
	}
	if !finished && !end {
		fmt.Println("进行中-")
	}
	if finished || end {
		fmt.Println("已完成或已结束！")
	}

}

func bit() {
	a := 0b00111100 // 60
	b := 0b00001101 // 13
	fmt.Printf("a&b=%.8b,a|b=%.8b,a^b=%.8b,\n", a&b, a|b, a^b)
	fmt.Printf("a<<2=%.8b,a>>2=%.8b\n", a<<2, a>>2)
}

func assignment() {
	var a = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)
}

func others() {
	var a = 1
	var ptr = &a
	fmt.Printf("a的地址：%p -> %d\n", ptr, a)
}
