package marshaller

import "go/types"

// Marshaller 序列化器
type Marshaller interface {
	// Marshal 序列化
	Marshal(params types.Array, payload types.Array) (string, error)

	// Unmarshal 反序列化
	Unmarshal(params string, payload types.Array) (types.Array, error)
}
