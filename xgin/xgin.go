package xgin

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/list":
		for k, v := range req.Header {
			fmt.Fprintf(w, "engine-Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func Xgin() {
	engine := new(Engine)
	/*
		// http.ListenAndServe 第二个参数Handler类型
		// Handler是一个接口，需要实现方法 ServeHTTP ，
		// 也就是说，只要传入任何实现了 ServerHTTP 接口的实例，
		// 所有的HTTP请求，就都交给了该实例处理了所有的HTTP请求，就都交给了该实例处理了
		// type Handler interface {
		// 	ServeHTTP(ResponseWriter, *Request)
		// }
	*/
	log.Fatal(http.ListenAndServe(":8081", engine))
}
