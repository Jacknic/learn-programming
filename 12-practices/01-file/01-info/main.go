package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	flag.Parse()
	args := flag.Args()
	filePath := "README.md"
	if len(args) > 0 {
		filePath = args[0]
	}
	file, err := os.Open(filePath)
	if file != nil {
		defer func() {
			_ = file.Close()
		}()
	}
	if err != nil {
		debug.PrintStack()
		message := fmt.Sprintf("打开文件失败！%s", filePath)
		panic(message)
	}
	stat, err := file.Stat()
	if err != nil {
		debug.PrintStack()
		panic("读取文件信息失败！")
	}
	fmt.Printf("文件大小：%s\n创建时间：%d\n文件权限：%O", fmtSize(stat.Size()), stat.ModTime().Unix(), stat.Mode())
}

//units 容量单位
const units = "BKMGTE"

func fmtSize(size int64) string {
	calcSize := float64(size)
	i := 0
	const kb = 1024
	for calcSize > kb && i < len(units) {
		calcSize /= kb
		i += 1
	}
	return fmt.Sprintf("%.2f%c", calcSize, units[i])
}
