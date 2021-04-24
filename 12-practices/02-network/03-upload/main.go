package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"runtime/debug"
)

const htmlUpload = `
<html>
<head>
	<title>上传文件</title>
</title>
<body>
	<form enctype="multipart/form-data" action="" method="post">
		<input type="file" name="upload" />
		<input type="submit" value="上传" />
	</form>
</body>
</html>
`

// 文件上传请求处理
func upload(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {
		_ = r.ParseMultipartForm(32 << 20)
		uploadFile, header, err := r.FormFile("upload")
		if err != nil {
			debug.PrintStack()
			log.Fatal("解析文件错误", err)
		}
		defer func(f multipart.File) { _ = f.Close() }(uploadFile)
		_ = os.Mkdir("out", 0666)
		filePath := "out/" + header.Filename
		saveFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal("创建存储文件失败：", err)
			return
		}
		defer saveFile.Close()
		_, err = io.Copy(saveFile, uploadFile)
		_, _ = io.WriteString(w, fmt.Sprintf("文件信息：%v", header))
		wd, _ := os.Getwd()
		fmt.Println("文件存放位置：", wd+"/"+filePath)
	} else {
		t, _ := template.New("upload").Parse(htmlUpload)
		_ = t.Execute(w, nil)
	}
}

// 文件上传处理
func main() {
	http.HandleFunc("/", upload)
	fmt.Println("访问地址 http://localhost:9090/")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("监听服务启动失败：", err)
		return
	}
}
