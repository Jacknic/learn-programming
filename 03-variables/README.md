# 变量

## 变量声明

- 不指定变量

  在基础数据类型未赋值时，为类型默认值；派生类型，默认值为nil

    ```go
  var numberNine int
  numberNine = 9
    ```

- 类型自动推断声明变量

    ```go
    numberNine := 9
    ```

- 多变量声明

  ```go
  var vname1,vname2,vname3 type
  vname1,vname2,vname3 = v1,v2,v3
  var vname1,vname2,vname3 = v1,v2,v3
  vname1,vname2,vname3 := v1,v2,v3
  
  var(
  vname1 type1
  vname2 type2
  )
  ```

## 忽略值

Go 中使用 `_` 忽略不需要的表达式值

```go
_,b := 5,7
```

