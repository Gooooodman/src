package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return

}
func main() {
	getImg("http://a.hiphotos.baidu.com/zhidao/wh%3D450%2C600/sign=4053015da7efce1bea7ec0ce9a61dfe8/f31fbe096b63f624bbdc7ba88744ebf81b4ca39c.jpg")
}

//该片段来自于http://outofmemory.cn
