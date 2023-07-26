package configer

type APIConfigure interface {
	// GetAPI 获取 API
	GetAPI(apiCode string) string
}
