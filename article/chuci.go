package article

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoadChuci() {
	dataUrl := "http://www.guoxue.com/jibu/chuci/content/cc_20.htm"
	testStr := "你好！"
	fmt.Println(testStr)
	resp, err := http.Get(dataUrl)
	if err != nil {

	}

	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
