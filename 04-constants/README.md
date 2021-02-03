# 常量

常量在程序运行时值已确定且不会被修改。Go 语言常量只可以是布尔、数值和字符串

- 普通定义

  ```go
  const NINE = 9
  ```

- 多值定义

  ```go
  const seven,eight,nine = 7,8,9
  ```

- 统一定义

  ```go
  const (
      Unknown = 0
      Female = 1
      Male = 2
  )
  ```

  ```go
  const (
  	Unknown = iota
  	Female
  	Male
  )
  ```

  

