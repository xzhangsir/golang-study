package cache

import (
	"fmt"
	"golang-study/cache/consistenthash"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_cache/"
	defaultReplicas = 50
)

type HttpPool struct {
	self     string
	basePath string
	mu       sync.Mutex
	// 一致性哈希算法的 Map，用来根据具体的 key 选择节点
	peers       *consistenthash.Map
	httpGetters map[string]*httpGetter
}

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (h *HttpPool) Log(format string, v ...interface{}) {
	log.Printf("[server %s] %s", h.self, fmt.Sprintf(format, v...))
}

func (h *HttpPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, h.basePath) {
		panic("意外的路径")
	}
	h.Log("%s %s", r.Method, r.URL.Path)
	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(r.URL.Path[len(h.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}

func (h *HttpPool) Set(peers ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.peers = consistenthash.New(defaultReplicas, nil)
	h.peers.Add(peers...)
	h.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		h.httpGetters[peer] = &httpGetter{baseURL: peer + h.basePath}
	}
}
func (h *HttpPool) Picker(key string) (PeerGetter, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if peer := h.peers.Get(key); peer != "" && peer != h.self {
		h.Log("Pick peer %s", peer)
		return h.httpGetters[peer], true
	}
	return nil, false
}

// HTTP 客户端类
type httpGetter struct {
	baseURL string
}

func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(group),
		url.QueryEscape(key),
	)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	return bytes, nil

}

// 编译时检查 *httpGetter 类型是否确实实现了 PeerGetter 接口
var _ PeerGetter = (*httpGetter)(nil)

// 确保HttpPool实现了接口PeerPicker
var _ PeerPicker = (*HttpPool)(nil) // 也可写成  var _PerrPicker = &Http{}
