package configer

type ClientConfigure interface {
	// GetValue 获取配置值
	GetValue(key string) string
}
