package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"strconv"
	"time"
)

/*响应信息结构体*/
type fmRoot struct {
	R                   int
	Is_show_quick_start int
	Song                []fmSong
}

//音乐的信息
type fmSong struct {
	Album       string
	Picture     string
	Ssid        string
	Artist      string
	Url         string
	Company     string
	Title       string
	Rating_avg  int
	Length      int64
	Subtype     string
	Public_time string
	Sid         string
	Aid         string
	Sha256      string
	Kbps        string
	Albumtitle  string
	Like        int
}

/*统计信息*/
type Sinfo struct {
	etime     float32
	totalCot  int
	totalMer  int
	succeeCot int
}

//常量 不可变
const (
	dir = "E:\\"
)

//抓取音乐列表
func smain() []byte {
	client := &http.Client{}
	// 获取的 url 网址http://douban.fm/j/mine/playlist?type=n&sid=347955&pt=2.9&channel=-2&pb=64&from=mainsite&r=1d815d7ebf
	// 每次获取都是不同的清单
	req, _ := http.NewRequest("GET", "http://douban.fm/j/mine/playlist?type=n&sid=347955&pt=2.9&channel=-2&pb=64&from=mainsite&r=1d815d7ebf", nil)
	//Add 是加的一些网页的头信息
	req.Header.Add("Accept:", "*/*")
	//req.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Add("Cookie", "openExpPan=Y; bid=\"8YicOUp+Kx0\"; dbcl2=\"56345753:wqSu+cIqVy8\"; fmNlogin=\"y\"; __utma=58778424.1673168777.1384438525.1385773681.1385773681.15; __utmb=58778424.1.9.1385790386347; __utmc=58778424; __utmz=58778424.1384438525.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)")
	req.Header.Add("Host", "douban.fm")
	req.Header.Add("Referer", "http://douban.fm/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36")
	resp, err := client.Do(req)
	// 返回网页的状态码 200 是成功的意思
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return body
	} else {
		log.Fatal(err)
		os.Exit(1)
	}
	return nil
}

//看程序从main开始看,看其逻辑
func main() {
	fmt.Println(`        ##############################################
        #############   下载豆瓣音乐     #############        
        ##############################################
                        保存路径：E盘(次数为10)
        `)
	fmt.Println("Spilder start................!")
	startTime := time.Now().UnixNano()
	sinfo := Sinfo{}
	sum := 1
	//这里访问10次  http://douban.fm/j/mine/playlist?type=n&sid=347955&pt=2.9&channel=-2&pb=64&from=mainsite&r=1d815d7ebf
	for sum < 11 {
		data := smain()
		if data != nil {
			root := &fmRoot{0, 0, nil}
			//data := string(data)
			err := json.Unmarshal(data, root)
			if err != nil {
				fmt.Println("Start Down File Err....")
				log.Fatal(err)
			} else {
				fmt.Printf("第 %d 次下载..\n", sum)
				//这里是把获取的信息放在一个struct中 方便调用它的key(键)获取vaule(值)
				crawMusic(root.Song, &sinfo)
				//log.Fatal(err)
			}
		}
		sum += 1
	}

	sinfo.etime = float32(time.Now().UnixNano()-startTime) / 1e9
	fmt.Printf("success cot %d  \n", sinfo.totalCot)
	fmt.Printf("exceute time %.3fs\n", sinfo.etime)
}

/*
*抓取
 */
func crawMusic(song []fmSong, sinfo *Sinfo) {
	sinfo.totalCot = sinfo.totalCot + len(song)
	for _, v := range song {
		if !FileExist(dir+v.Title+".mp4") && v.Url != "" {
			//log.Println(v.Url + " " + v.Title + "  " + strconv.FormatInt(v.Length, 10))
			log.Println("开始下载歌曲: " + v.Title)
			downMusic(v, sinfo)
		}
	}
}

//这个函数就是得到了 mp4的url 进行读取返回内容
func downFile(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Add("Cookie", "openExpPan=Y; bid=\"8YicOUp+Kx0\"; dbcl2=\"56345753:wqSu+cIqVy8\"; fmNlogin=\"y\"; __utma=58778424.1673168777.1384438525.1385773681.1385773681.15; __utmb=58778424.1.9.1385790386347; __utmc=58778424; __utmz=58778424.1384438525.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)")
	req.Header.Add("Host", "mr3.douban.com")
	req.Header.Add("Referer", "http://douban.fm/")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.76 Safari/537.36")
	resp, _ := client.Do(req)

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return body
	} else {
		log.Println("error -->" + url)
	}
	return nil
}

/*
*下载文件
 */
func downMusic(song fmSong, sinfo *Sinfo) {

	file := dir + song.Title + ".mp4"
	dts := downFile(song.Url)
	if dts != nil {
		defer func() {
			f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			defer f.Close()
			//将读取到的信息保存在本地文件中
			ioutil.WriteFile(file, dts, os.ModePerm)
			//fmt.Println(file)
			sinfo.succeeCot++
		}()
	} else {
		log.Println("save file error")
	}
}

/*
*保存歌词详细信息
 */
func FileInfo() {
	//TODO
}

/*
*文件是否存在
 */
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
