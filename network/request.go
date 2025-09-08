package network

// Request 定义了一个 HTTP 请求的结构体
type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Params  map[string]string
	Data    map[string]any
	Json    any
}
