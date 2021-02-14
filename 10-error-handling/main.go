package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	if result, err := div(10, 0); err != nil {
		fmt.Println("计算过程发生错误:", reflect.TypeOf(err), err.Error())
	} else {
		fmt.Println(`计算结果为:`, result)
	}
	files := []string{"NotFound.txt", "main.go"}
	for _, fileName := range files {
		if file, err := os.Open(fileName); err != nil {
			// 类型转换
			if pathError, ok := err.(*os.PathError); ok {
				fmt.Println("打开失败，文件路径有误：", pathError.Path)
			}
		} else {
			reader := bufio.NewReader(file)
			writer := bufio.NewWriter(os.Stdout)
			written, err := io.Copy(writer, reader)
			if err != nil {
				fmt.Println("写入数据失败：", err)
			} else {
				fmt.Printf("共计写入:%d\n", written)
			}
		}
	}

}

type DivError struct {
	message string
}

func (err DivError) Error() string {
	return err.message
}

func div(a, b int) (result int, err error) {
	if b == 0 {
		err = &DivError{
			"除数不能为0",
		}
		return -1, err
	}
	return a / b, err
}
