package http

import (
	"fmt"
	"mouse/command"
	"net/http"
	_ "strings"
)

func HttpServer() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/json", jsonHandler)
	_ = http.ListenAndServe("localhost:19527", nil)
}

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	command.Lissajous(writer)
}

func jsonHandler(writer http.ResponseWriter, request *http.Request) {
	content := "你好，测试下GO的web速度"
	_, _ = writer.Write([]byte(content))
}
