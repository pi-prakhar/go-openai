package api

import (
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	mux.HandleFunc("/test/api", HandleOpenAITest)
	mux.HandleFunc("/api/chat/send", HandleSendChatMessage)
	mux.HandleFunc("/api/chat/messages", HandleGetMessages)
	return mux
}
