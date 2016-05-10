package main

import (
	//	"flag"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	//"path/filepath"
	"regexp"
	"strings"
)

func walkFile(dir string, filenames chan<- string) {
	r, err := regexp.Compile(`.*\.txt$`)
	for _, entry := range dirents(dir) {
		// if entry.IsDir() {
		// 	fmt.Println(entry.Name())
		// 	subdir := filepath.Join(dir, entry.Name())
		// 	fmt.Sprint(subdir,"\\",)
		// 	walkFile(subdir, filenames)
		// } else {
		if !entry.IsDir() {
			if err != nil {
				fmt.Println(err)
			}
			if r.FindString(entry.Name()) != "" {
				filenames <- entry.Name()
			}
		}
	}
}

func desc() {
	fmt.Println("                ##############################################")
	fmt.Println("                #############  批量修改后缀名  ###############")
	fmt.Println("                #############  当前目录输入 .  ###############")
	fmt.Println("                ##############################################")
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func bk() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	desc()
	fmt.Printf("请输入文件夹位置(如：f:\\good): ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("文件夹位置不正确..")
		bk()
		return
	}
	line = strings.TrimSpace(line)
	if line == "" {
		fmt.Println("文件夹位置未填..")
		bk()
		return
	}
	fmt.Printf("请输入被改变的文件后缀(如：txt): ")
	sour_suff := bufio.NewReader(os.Stdin)
	sorc, err := sour_suff.ReadString('\n')
	if err != nil {
		fmt.Println("被改变的文件后缀错误: ", err)
		bk()
		return
	}
	sorc = strings.TrimSpace(sorc)
	fmt.Printf("请输入改变后的文件后缀(如：cpp): ")
	now_suff := bufio.NewReader(os.Stdin)
	now, err := now_suff.ReadString('\n')
	if err != nil {
		fmt.Println("改变后的文件后缀错误: ", err)
		return
	}
	now = strings.TrimSpace(now)
	filenames := make(chan string)
	//fmt.Println(sorc, now)
	go func() {
		walkFile(line, filenames)
		//}
		close(filenames)
	}()

loop:
	for {
		select {
		case f, ok := <-filenames:
			if !ok {
				break loop // fileSizes was closed
			}
			top := strings.Split(f, ".")
			end := fmt.Sprint(top[0], ".", now)
			os.Rename(f, end)

		}
	}
	fmt.Println("更改完成...按任意键退出!!")
	bk()
}
