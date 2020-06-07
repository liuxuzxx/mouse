package goroutines

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	host string
	port string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("https://" + h.host + ":" + h.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func StartProxy() {
	response, _ := http.Get("https://open.e.189.cn/auth/sdkcodeinfo.do")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(body))
}