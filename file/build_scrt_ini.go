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
	// fmt.Println(*FileName)
	// fmt.Println(*Name)
	// fmt.Println(*Conffile)
	// f, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// 	panic(err)
	// }
	// fmt.Println(string(f))
	//fmt.Println("########################################################################")
	filelist := []string{}

	b := Exist(*FileName)
	if !b {
		fmt.Println("ini文件不存在.")
		Use()
		return
	}
	b = Exist(*Conffile)
	if !b {
		fmt.Println("hostname文件不存在.")
		Use()
		return
	}
	context := read3(*FileName)
	// context = strings.Replace(context, "120.55.164.160", "139.196.179.114", -1)
	// //fmt.Println(context)
	// file1, _ := os.Create("al_gwlm_gs_9011_139-196-179-114.ini")
	f, err := os.Open(*Conffile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		line = strings.TrimSpace(line)
		if err != nil || io.EOF == err {
			break
		}
		filelist = append(filelist, line)
	}
	// fmt.Println(filelist)
	// fmt.Println(len(filelist))
	// for i, n := range filelist {
	// 	fmt.Println(i, "----->", n)
	// }
	// return
	//chs := make(chan int, len(filelist))
	for _, f := range filelist {
		//chs[i] = make(chan int)
		//go func(f string, c chan int) {
		//f = strings.Replace(f, "\n", "", 2)
		//n := strings.Replace(f, "\n", ".ini", 1)
		//n = n + ".ini"
		//fmt.Println(n)
		newfile := fmt.Sprintf("%s.ini", f)
		//fmt.Println(strings.Replace(newfile, "\n", "", 1))
		//fmt.Println(newfile)
		a := strings.Split(f, "_")
		newip := strings.Replace(a[4], "-", ".", -1)
		//return
		//newip := "000.000.000.000"
		WriteFile(context, newfile, *Name, newip)
		//}(f, chs[i])
	}
	//WriteFile(context, "default.ini", "root", "1.1.1.1")
	// defer file1.Close()
	// n, err := io.WriteString(file1, context)
	// if err != nil {
	// 	fmt.Println(n)
	// }
	// for _, ch := range chs {
	// 	fmt.Println(<-ch)
	// }

}

func WriteFile(context string, newfile string, newname string, newip string) {
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

}
