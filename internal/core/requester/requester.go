package requester

import "go/types"

// Requester 请求器
type Requester interface {
	// SendGet 发送GET请求
	SendGet(apiCode string, businessParams types.Array) (types.Array, error)

	// SendPost 发送POST请求
	SendPost(apiCode string, businessParams types.Array) (types.Array, error)
}
