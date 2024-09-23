package cache

// 根据传入的key选择相应节点的PeerGetter
type PeerPicker interface {
	Picker(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
