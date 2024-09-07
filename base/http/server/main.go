package main

import (
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	str := "hello"
	w.Write([]byte(str))
}
func getUrl(w http.ResponseWriter, r *http.Request) {
	// 处理客户端的请求
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	// 获取GET请求的参数 URL中携带的参数
	queryParam := r.URL.Query()
	fmt.Println(queryParam)
	fmt.Println(queryParam.Get("user"))
	// 获取POST请求的参数
	// fmt.Println(ioutil.ReadAll(r.Body))
	// 发送数据给客户端
	w.Write([]byte("ok"))
}
func postUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.body err:%v", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/userinfo", getUrl)
	http.HandleFunc("/list", postUrl)
	err := http.ListenAndServe("127.0.0.1:20000", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
