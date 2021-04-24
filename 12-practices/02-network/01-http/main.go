package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 发送http请求
func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("请求重定向操作")
			return nil
		},
	}
	resp, err := client.Get("https://baidu.com")
	if err != nil {
		panic("发送请求失败：" + err.Error())
	}
	_, err = io.Copy(os.Stdout, resp.Body)
}
