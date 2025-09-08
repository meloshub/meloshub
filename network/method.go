package network

var defaultSession = NewSession()

// Get 发送一个 GET 请求
func Get(url string, params map[string]string) (*Response, error) {
	return defaultSession.Get(url, params)
}

// PostJSON 发送一个 POST 请求，内容为 JSON
func PostJSON(url string, jsonData any) (*Response, error) {
	return defaultSession.PostJSON(url, jsonData)
}

// PostForm 发送一个 POST 请求，内容为FormData
func PostForm(url string, data map[string]any) (*Response, error) {
	return defaultSession.PostForm(url, data)
}