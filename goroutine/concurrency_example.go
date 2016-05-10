package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		//"http://www.g5od.com",
		"http://www.sina.com.cn",
		"http://www.baidu.com",
		"http://www.qq.com",
	}
	results := make(chan HomePageSize)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}
	var biggest HomePageSize

	//	fmt.Println(biggest)
	for range urls {
		result := <-results
		//fmt.Println(result)
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println(biggest)
	fmt.Println("The biggest home page:", biggest.URL)
}
