package network

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response HTTP 响应的结构体
type Response struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       []byte
	Request    *http.Request
}

// HTTPError 自定义错误类型，用于表示非 2xx 的 HTTP 状态码
type HTTPError struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (e *HTTPError) Error() string {
	bodySample := string(e.Body)
	if len(bodySample) > 100 {
		bodySample = bodySample[:100] + "..."
	}
	return fmt.Sprintf("request failed with status %s, body: %s", e.Status, bodySample)
}

// Text 返回响应体的字符串表示
func (r *Response) Text() string {
	return string(r.Body)
}

// Bytes 返回原始的响应体字节切片
func (r *Response) Bytes() []byte {
	return r.Body
}

// JSON 解析响应体为 JSON，并将结果存储在 v 指向的变量中
func (r *Response) JSON(v any) error {
	if len(r.Body) == 0 {
		return fmt.Errorf("response body is empty")
	}
	return json.Unmarshal(r.Body, v)
}

// IsSuccess 检查 HTTP 状态码是否为 2xx。如果不是，则返回一个 HTTPError。
func (r *Response) IsSuccess() error {
	if r.StatusCode >= 200 && r.StatusCode <= 299 {
		return nil
	}
	return &HTTPError{
		StatusCode: r.StatusCode,
		Status:     r.Status,
		Body:       r.Body,
	}
}
