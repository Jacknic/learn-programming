# 基础语法

## Go 标记

Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。如以下 GO 语句由 6 个标记组成：

```go
fmt.Println("Hello, World!")
```

## 行分隔符

在 Go 程序中，一行代表一个语句结束。如多个语句同行需用 `;` 分隔

```go
fmt.Println("Hello, World!")
num1 := 1;num2 :=2
```

## 字符串连接

Go 语言使用 `+` 拼接字符串

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello " + "world")
}
```

## 关键字

Go 中有 25 个关键字或保留字，不可用作变量名

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

36 个预定义标识符：

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |
