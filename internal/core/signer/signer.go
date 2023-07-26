package signer

import "go/types"

// Signer 签名器
type Signer interface {
	// CalcSign 计算签名
	calcSign(array types.Array) string
}
