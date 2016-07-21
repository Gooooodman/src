package main

// import (
// 	"fmt"
// 	"github.com/opesun/goquery"
// 	"strings"
// )

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/opesun/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	//"time"
)

const (
	Prefix = "http://imgsrc.baidu.com"
	Suffix = "jpg"
)

var url = flag.String("url", "http://tieba.baidu.com/p/4676296099", "separator")

func main() {
	flag.Parse()
	p, err := goquery.ParseUrl(*url)
	if err != nil {
		panic(err)
	} else {
		pTitle := p.Find("title").Text() //直接提取title的内容
		fmt.Println(pTitle)
		t := p.Find("").Attrs("src")
		//fmt.Println(p.Find("pre#line1"))
		// t := p.Find(".attribute-value a")
		ch := make([]chan bool, len(t))
		// for _, h := range t {
		// 	Select_img(h)
		// }
		for i, h := range t {
			ch[i] = make(chan bool)
			go Select_img(h, ch[i])
		}

		for _, c := range ch {
			<-c
		}
	}

}

func Select_img(url string, ch chan bool) {
	uls := strings.Split(url, "/")
	//fmt.Println(uls[0], string(uls[2]))
	head := fmt.Sprintf("%s//%s", uls[0], uls[2])
	//fmt.Println(head)
	if head == Prefix {
		getImg(url)
	}
	ch <- true
	close(ch)
}

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
