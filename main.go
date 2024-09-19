package main

import (
	"fmt"
	"golang-study/cache"
	"log"
	"net/http"
)

func main() {
	// base.Init()
	// ginweb.Init()
	testCache()
}

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
func testCache() {
	cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := cache.NewHttpPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
