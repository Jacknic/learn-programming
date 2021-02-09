package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 不使用标记，无法跳出外层循环
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}

	fmt.Println("进行输入(ctrl+D退出输入)：")
	buffer := make([]byte, 0)
	// 死循环
	for {
		reader := bufio.NewReader(os.Stdin)
		line, prefix, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				fmt.Println("输入结束")
			} else {
				fmt.Println("输入流错误：", err.Error())
			}
			break
		}
		buffer = append(buffer, line...)
		if !prefix {
			buffer = append(buffer, []byte("\n")...)
		}

	}
	_, _ = os.Stdout.Write(buffer)

	// range 循环
	for index, value := range "123456" {
		fmt.Printf("%d -> %c\n", index, value)
	}
}
