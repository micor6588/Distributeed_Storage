package objects

import "net/http"

// Handler 处理注册函数
func Handler(w http.ResponseWriter, r *http.Request) {
	// Handler函数首先检查HTTP请求的方法：如果是PUT，则调用put函数
	// 如果是Get方法则调用get函数
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
