package main

import (
	"fmt"
	"golang-study/xgin"
	"net/http"
)

func main() {
	// base.Init()
	// ginweb.Init()
	// testCache()
	testXgin()
}

/*  分布式缓存
// type String string

// func (d String) Len() int {
// 	return len(d)
// }

// // 分布式缓存
// func main() {
// 	keys := make([]string, 0)
// 	callback := func(key string, value lru.Value) {
// 		keys = append(keys, key)
// 	}
// 	l := lru.NewCache(int64(10), callback)
// 	l.Add("key1", String("123"))
// 	l.Add("k2", String("234"))
// 	l.Add("k3", String("345"))
// 	l.Add("k4", String("456"))
// 	l.ShowList()

// }
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// http://localhost:9999/_cache/scores/Tom

func createGroup() *cache.Group {
	return cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, gee *cache.Group) {
	peers := cache.NewHttpPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, gee *cache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

}

func testCache() {
	// cache.NewGroup("scores", 2<<10, cache.GetterFunc(
	// 	func(key string) ([]byte, error) {
	// 		log.Println("[SlowDB] search key", key)
	// 		if v, ok := db[key]; ok {
	// 			return []byte(v), nil
	// 		}
	// 		return nil, fmt.Errorf("%s not exist", key)
	// 	}))

	// addr := "localhost:9999"
	// peers := cache.NewHttpPool(addr)
	// log.Println("cache is running at", addr)
	// log.Fatal(http.ListenAndServe(addr, peers))
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "cache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], []string(addrs), gee)
}

// go run main.go -port=8001
// go run main.go -port=8002
// go run main.go -port=8003 -api=1
// http://localhost:9999/api?key=Tom
*/
// xgin框架开发
func testXgin() {
	r := xgin.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":8081")
}
