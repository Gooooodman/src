package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkFile(dir string) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkFile(subdir)
		} else {
			//filenames <- entry.Name()
			//fmt.Println("name: ", entry.Name())
			//expr,err := regexp.Compile(expr) ".*(.txt)$"
			r, err := regexp.Compile(`.*\.txt$`)
			if err != nil {
				fmt.Println(err)
			}
			if r.FindString(entry.Name()) != "" {
				fmt.Println(r.FindString(entry.Name()))
			}
		}
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	//filenames := make(chan string)
	for _, root := range roots {
		walkFile(root)
	}

}
