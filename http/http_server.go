package http

import (
	"fmt"
	"mouse/command"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/", httpHandler)
	_ = http.ListenAndServe("localhost:19527", nil)
}

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	command.Lissajous(writer)
}
