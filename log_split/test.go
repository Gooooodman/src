package main

import (
	"bufio"
	//"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func ReadLine(fileName string) (error, map[int]map[string]string) {
	f, err := os.Open(fileName)
	if err != nil {
		return err, nil
	}
	buf := bufio.NewReader(f)

	Binfo := make(map[int]map[string]string)
	num := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			//读完了之后 break
			if err == io.EOF {
				//return nil, Binfo
				break
			}
			return err, nil

		}
		line = strings.TrimSpace(line)
		//handler(line)
		r, _ := regexp.Compile("T\\d{2}:\\d{2}:\\d{2}.*")
		o := r.ReplaceAllString(line, "")
		r, _ = regexp.Compile("[[:space:]]")
		//fmt.Println(strings.Replace(r.ReplaceAllString(o, "&localtime="), "http://sdklog.tj.65.com/1.html?", "", 1))
		S_line := strings.Replace(r.ReplaceAllString(o, "&localtime="), "http://sdklog.tj.65.com/1.html?", "", 1)
		sp_one := strings.Split(S_line, "&")

		info, ok := Binfo[num]
		if !ok {
			info = make(map[string]string)
			Binfo[num] = info
			for _, v := range sp_one {
				v1 := strings.Split(v, "=")
				//fmt.Println(v1[1])
				info[v1[0]] = v1[1]
			}
			num++
		}

	}
	return nil, Binfo
}

func main() {

	// ReadLine("test2.log", func(out string) {
	// 	r, _ := regexp.Compile("T\\d{2}:\\d{2}:\\d{2}.*")
	// 	o := r.ReplaceAllString(out, "")
	// 	r, _ = regexp.Compile("[[:space:]]")

	// 	fmt.Println(strings.Replace(r.ReplaceAllString(o, "&localtime="), "http://sdklog.tj.65.com/1.html?", "", 1))
	// 	line := strings.Replace(r.ReplaceAllString(o, "&localtime="), "http://sdklog.tj.65.com/1.html?", "", 1)
	// 	sp_one := strings.Split(line, "&")
	// 	for _, v := range sp_one {
	// 		v1 := strings.Split(v, "=")
	// 		//fmt.Println(v1[1])
	// 		info[v1[0]] = v1[1]
	// 	}
	// 	Binfo[num] = info
	// 	num++
	// 	fmt.Println(Binfo)
	// })

	e, B := ReadLine("test2.log")
	if e == nil {
		fmt.Println(B)
	}
}
