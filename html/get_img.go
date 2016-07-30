package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	url = "http://tieba.baidu.com/p/4666577302"
)

func main() {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//context := string(pix)
	r, err := regexp.Compile("img class=\"BDE_Image\" src=\"(.+?\\.jpg)\"")
	if err != nil {
		fmt.Println(err)
	}
	imgs := r.FindAll(pix, -1)
	// fmt.Println(imgs)
	reg := regexp.MustCompile(`(.*)src=`)
	//rep := []byte("")

	var urls []string
	for _, i := range imgs {
		//fmt.Println(string(i))

		//strings.Replace(string(i), "")
		//fmt.Printf("%s\n", reg.ReplaceAllLiteral(i, rep))   //参数为[]byte
		u := reg.ReplaceAllLiteralString(string(i), "") //参数为string
		urls = append(urls, u)
	}
	fmt.Println(urls)
	//return
	ch := make([]chan bool, len(urls))
	// for _, h := range t {
	//  Select_img(h)
	// }
	for i, h := range urls {
		ch[i] = make(chan bool)
		go Select_img(h, ch[i])
		// go func(url string, ch chan bool) {
		// 	uls := strings.Split(url, "\"")    getImg 要去掉引号
		// 	//fmt.Println(uls[0])
		// 	getImg(string(uls[1]))
		// 	ch <- true
		// 	close(ch)
		// }(h, ch[i])
	}

	for _, c := range ch {
		<-c
	}
}

func Select_img(url string, ch chan bool) {
	//uls := strings.Split(url, "/")
	//fmt.Println(uls[0], string(uls[2]))
	// head := fmt.Sprintf("%s//%s", uls[0], uls[2])
	//fmt.Println(head)
	//if head == Prefix {
	//fmt.Println(url)
	uls := strings.Split(url, "\"")
	getImg(string(uls[1]))
	//getImg(url)
	//}
	ch <- true
	close(ch)
}

func getImg(url string) (n int64, err error) {
	//fmt.Println(url)
	//return
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
