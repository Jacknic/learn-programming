package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const htmlLogin = `
<html>
<head>
	<title>用户登录</title>
</head>
<body>
	<form action="" method="POST">
		用户名：<input type="text" name="username" />
		密码：<input type="text" name="password" />
		<input type="submit" value="登录" />
	</form>
</body>
</html>
`

// 用户登录
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("当前请求方法：", r.Method)
	if "POST" == r.Method {
		_ = r.ParseForm()
		username := r.Form["username"]
		password := r.Form["password"]
		fmt.Println("username:", username, "password:", password)
		message := fmt.Sprintf("username:%s,password:%s", username, password)
		_, _ = w.Write([]byte(message))
	} else {
		t := template.New("login")
		_, _ = t.Parse(htmlLogin)
		_ = t.Execute(w, nil)
	}
}

// 处理表单数据
func main() {
	http.HandleFunc("/login", login)
	fmt.Println("访问地址 http://localhost:9090/login")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("创建端口监听失败：", err)
	}
}
