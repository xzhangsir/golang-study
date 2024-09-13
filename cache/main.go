package main

import (
	"golang-study/cache/lru"
)

type String string

func (d String) Len() int {
	return len(d)
}

// 分布式缓存
func main() {
	keys := make([]string, 0)
	callback := func(key string, value lru.Value) {
		keys = append(keys, key)
	}
	l := lru.NewCache(int64(10), callback)
	l.Add("key1", String("123"))
	l.Add("k2", String("234"))
	l.Add("k3", String("345"))
	l.Add("k4", String("456"))
	l.ShowList()

}
