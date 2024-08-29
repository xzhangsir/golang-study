package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
	"strings"
)

/*func main(){
	// 发送get请求给服务端
	resq,err := http.Get("http://127.0.0.1:20000/userinfo?user=zx")
	if err != nil{
		fmt.Printf("Get failed ,err: %v\n",err)
		return
	}
	defer resq.Body.Close()
	// 接受服务端的返回
	body,err := ioutil.ReadAll(resq.Body)
	if err != nil{
		fmt.Printf("read from err:%v",err)
		return
	}
	fmt.Printf(string(body))
}*/


// 对GET请求  URL中的参数 进行 encode  使用net/url库

/*func main(){
	apiUrl := "http://127.0.0.1:20000/userinfo"
	data := url.Values{}
	data.Set("user","小星星")
	u,err := url.ParseRequestURI(apiUrl)
	if err != nil{
		fmt.Printf("parse url err %v\n",err)
		return
	}
	u.RawQuery = data.Encode() //url encode

	resp, err := http.Get(u.String())  //和下面的写法类似
	// req,err := http.NewRequest("Get",u.String(),nil)  //创建请求对象
	// resp,err := http.DefaultClient.Do(req) //发请求
	

	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

}
*/

// 发送post请求
func main(){
	url := "http://127.0.0.1:20000/list"
	// 表单
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=小王子&age=18"

	// JSON
	contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}