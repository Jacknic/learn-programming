package main

import (
	"encoding/json"
	"fmt"
)

//Book 书籍信息 Title 标题
type Book struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	IsPublished bool     `json:"is_published"`
	Price       float32  `json:"price"`
}

//main JSON 的简单读取
func main() {
	goBook := Book{
		"Go语音编程",
		[]string{"one", "two", "three"},
		"中华书局",
		true,
		9.99,
	}
	jsonBytes, err := json.Marshal(goBook)
	if err != nil {
		panic("序列化JSON对象失败")
	}
	fmt.Println(string(jsonBytes))
	jsonText := `{"title":"Go语音编程 第二版","authors":["one","two","three"],"publisher":"中华书局","is_published":true,"price":9.99}`
	var bookInfo Book
	if err := json.Unmarshal([]byte(jsonText), &bookInfo); err != nil {
		panic("实例化JSON数据对象失败")
	}
	bytes, err := json.Marshal(bookInfo)
	println(string(bytes))
}
