package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

// Session 定义了一个带有 Cookie 管理功能的 HTTP 会话
type Session struct {
	Client *http.Client
}

// newResponse 从 http.Response 创建一个自定义的 Response 对象
func newResponse(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header,
		Body:       body,
		Request:    resp.Request,
	}, nil
}

// NewSession 创建一个带 Cookie 管理功能的新会话
func NewSession() *Session {
	jar, _ := cookiejar.New(nil)
	return &Session{
		Client: &http.Client{
			Jar:     jar,
			Timeout: 30 * time.Second,
		},
	}
}

// Do 发送一个 HTTP 请求并返回响应
func (s *Session) Do(req *Request) (*Response, error) {
	var bodyReader io.Reader
	if req.Json != nil {
		jsonData, err := json.Marshal(req.Json)
		if err != nil {
			return nil, fmt.Errorf("json marshal failed: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
		if req.Headers == nil {
			req.Headers = make(map[string]string)
		}
		if _, ok := req.Headers["Content-Type"]; !ok {
			req.Headers["Content-Type"] = "application/json"
		}
	} else if req.Data != nil {
		formData := url.Values{}
		for k, v := range req.Data {
			formData.Set(k, fmt.Sprintf("%v", v))
		}
		bodyReader = strings.NewReader(formData.Encode())
		if req.Headers == nil {
			req.Headers = make(map[string]string)
		}
		if _, ok := req.Headers["Content-Type"]; !ok {
			req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
		}
	}

	targetURL, err := url.Parse(req.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}
	if req.Params != nil {
		q := targetURL.Query()
		for k, v := range req.Params {
			q.Set(k, v)
		}
		targetURL.RawQuery = q.Encode()
	}

	httpReq, err := http.NewRequest(req.Method, targetURL.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}
	httpResp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return newResponse(httpResp)
}

func (s *Session) SetCookieFromMap(cookies map[string]string, targetURL string) error {
	u, err := url.Parse(targetURL)
	if err != nil {
		return fmt.Errorf("failed to parse target URL: %w", err)
	}

	// 将 map 转换为 []*http.Cookie 切片
	var cookieList []*http.Cookie
	for name, value := range cookies {
		cookie := &http.Cookie{
			Name:  name,
			Value: value,
			Path:  "/",
		}
		cookieList = append(cookieList, cookie)
	}

	s.Client.Jar.SetCookies(u, cookieList)

	return nil
}

// Get 发送GET请求
func (s *Session) Get(url string, params map[string]string) (*Response, error) {
	req := &Request{
		Method: "GET",
		URL:    url,
		Params: params,
	}
	return s.Do(req)
}

// PostForm 发送POST表单请求
func (s *Session) PostForm(url string, data map[string]any) (*Response, error) {
	req := &Request{
		Method: "POST",
		URL:    url,
		Data:   data,
	}
	return s.Do(req)
}

// PostJSON 发送POST JSON请求
func (s *Session) PostJSON(url string, jsonData any) (*Response, error) {
	req := &Request{
		Method: "POST",
		URL:    url,
		Json:   jsonData,
	}
	return s.Do(req)
}
