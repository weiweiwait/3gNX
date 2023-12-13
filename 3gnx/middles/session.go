package middles

import (
	"net/http"
)

func GetSessionId(w http.ResponseWriter, r *http.Request) string {
	session, err := r.Cookie("Cookie")
	if err != nil {
		// 处理获取会话时的错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "错误"
	}

	// 从会话 cookie 中获取会话标识符
	sessionID := session.Value
	return sessionID

}
