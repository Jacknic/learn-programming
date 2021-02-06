package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// 布尔类型
	const OK = true
	const ENABLED = false
	fmt.Println(OK, ENABLED)

	// 数字类型 int
	maxFloat32 := math.MaxFloat32
	maxInt64 := math.MaxInt64
	maxInt8 := math.MaxInt8
	fmt.Println(maxFloat32)
	fmt.Println(maxInt64 >> 60)
	fmt.Println(maxInt64 & 0b11)

	// 字符串类型
	const message = "send it to my phone"
	fmt.Println(strings.Split(message, " ")[4])

	// 指针类型
	fmt.Println("指针值变化：")
	i32 := &maxInt8
	fmt.Println("原始值", maxInt8)
	changeValue32(maxInt8)
	fmt.Println("修改值", maxInt8)
	fmt.Println("原始值", *i32)
	changeValue(i32)
	fmt.Println("修改值", *i32)

	fmt.Printf("%v->%d\n", i32, *i32)

	// 数组类型
	var numbers0 = [5]float32{10.0, 9.0, 8, 7, 6}
	numbers1 := [...]float32{20.0, 9.0, 8, 7, 6}
	numbers2 := [...]float32{1: 2, 9: 4}
	matrix := [...][3]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(numbers0, numbers1, numbers2)
	fmt.Println(numbers0[2])
	fmt.Println(matrix)

	// 结构化数据

	userInfo := UserInfo{"Jacknic", 22, 1}
	fmt.Println("用户数据", userInfo)
	fmt.Println(userInfo.name)

	// 通道
	chanInt := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		chanInt <- i
	}
	fmt.Println(<-chanInt)
	fmt.Println(<-chanInt)
	//for i := range chanInt {
	//	fmt.Println(i)
	//}

	// 函数类型
	method := sayHi
	method()

	// 切片类型
	varSlice := make([]int, 10)
	for i := range varSlice {
		varSlice[i] = i
	}
	fmt.Println(varSlice)
	fmt.Println(varSlice[2:6])

	// 接口类型
	userInfo.Run()
	var runnable = userInfo.Run
	typeOf := reflect.TypeOf(runnable)
	fmt.Println("类型名称:", typeOf.Name())
	runnable()

	// Map 类型
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Japan"] = "Tokyo"
	for k, e := range countryCapitalMap {
		fmt.Println("country:", k, "capital:", e)
	}
	China := "China"
	capital, ok := countryCapitalMap[China]
	if ok {
		fmt.Printf("%s -> %s", China, capital)
	} else {
		fmt.Printf("China's capital not found")
	}
}

//changeValue 将数值改为32
func changeValue(prt *int) {
	*prt = 32
}

func changeValue32(value int) {
	value = 32
}

type UserInfo struct {
	name   string
	age    int
	gender byte
}

//Runnable 定义接口
type Runnable interface {
	Run()
}

//Run 实现接口
func (userInfo UserInfo) Run() {
	fmt.Println("打印用户名：", userInfo.name)
}

//sayHi 打印hi
func sayHi() {
	fmt.Println("Hi!")
}
