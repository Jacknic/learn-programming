package main

import (
	"fmt"
	"math"
	"os"
	osx "os"
)

// UserInfo 用户信息
type UserInfo struct {
	Name   string
	Age    int
	Gender byte
	Weight float64
}

func main() {
	// 输出语句
	const message = "Hello " + "world"
	fmt.Println(message, len(message))
	loop()

	var result = div(1, 0)
	println(result)

	readFile()
	showCollections()
	fmt.Println(osx.Environ()[0:2])
}

//loop 循环与条件判断
func loop() {
	for i := 1; i < 10; i++ {
		if i >= 5 {
			fmt.Println("跳出循环:", i)
			break
		} else if i == 4 {
			fmt.Printf("跳过该数:%d\n", i)
			continue
		}
		switch i {
		case 1:
			fmt.Println("一")
			break
		case 2:
			fmt.Println("二")
			break
		default:
			fmt.Printf("输出原数：%d\n", i)
		}
	}
}

//showCollections 集合类型数据
func showCollections() {
	userInfoList := []UserInfo{
		{"Jack", 18, 0, math.Round(50)},
		{"Tome", 7, 0, math.Round(50)},
		{Name: "Jerry", Age: 6, Gender: 0, Weight: math.Round(50)},
	}
	userInfoMap := make(map[string]int)
	for i := 0; i < len(userInfoList); i++ {
		info := userInfoList[i]
		userInfoMap[info.Name] = info.Age
	}
	fmt.Println(userInfoList)
	fmt.Println(userInfoMap)
}

//readFile 读取文件
func readFile() {
	readme, err := os.Open("README.md")
	if readme != nil {
		defer readme.Close()
		buff := make([]byte, 14)
		count, _ := readme.Read(buff)
		// fmt.Println(string(count) + "--" + err)
		_, _ = os.Stdout.Write(buff[0:count])
		_, _ = os.Stdout.WriteString("\nDone\n")
	}
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
}

//div 两数相除
func div(num1 int, num2 int) int {
	// 异常捕获
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获异常：", err)
		}
	}()
	return num1 / num2
}
