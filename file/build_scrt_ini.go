package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

var FileName = flag.String("file", "default.ini", "input source file")
var Conffile = flag.String("conf", "default.conf", "hostname  file")
var Name = flag.String("name", "lupuxiao", "modify ini  name")

func Use() {
	fmt.Println("-file [default.ini] -conf [default.ini] -name [lupuxiao]")
}

func main() {
	flag.Parse()
	filelist := []string{}

	b := Exist(*FileName)
	if !b {
		fmt.Println("初始default.ini文件不存在.")
		Use()
		return
	}
	b = Exist(*Conffile)
	if !b {
		fmt.Println("初始default.conf文件不存在.")
		Use()
		return
	}
	context := read3(*FileName)
	f, err := os.Open(*Conffile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		line = strings.TrimSpace(line)   // 去掉前后端所有空白
		if err != nil || io.EOF == err {
			break
		}
		filelist = append(filelist, line)
	}
	chs := make([]chan int, len(filelist))
	for i, f := range filelist {
		chs[i] = make(chan int)
		//go func(f string, c chan int) {
		newfile := fmt.Sprintf("%s.ini", f)
		a := strings.Split(f, "_")
		newip := strings.Replace(a[4], "-", ".", -1)
		go WriteFile(context, newfile, *Name, newip, chs[i])
		//}(f, chs[i])
	}
	for _, ch := range chs {
		<-ch
	}

}

func WriteFile(context string, newfile string, newname string, newip string, c chan int) {
	r, _ := regexp.Compile("S:\"Hostname\"=.*")
	//context = strings.Replace(context, "120.55.164.160", newip, -1)
	Hostname := fmt.Sprintf("S:\"Hostname\"=%s", newip)

	context1 := r.ReplaceAllString(context, Hostname)
	r1, _ := regexp.Compile("S:\"Username\"=.*")
	Username := fmt.Sprintf("S:\"Username\"=%s", newname)
	newcontext := r1.ReplaceAllString(context1, Username)
	// file, _ := os.Create(newfile)
	// defer file.Close()
	err := ioutil.WriteFile(newfile, []byte(newcontext), 0666)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s write Ok..\n", newfile)
	c <- 1
	close(c)
}
