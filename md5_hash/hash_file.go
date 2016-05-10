package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

/*计算文件的md5 sha1*/

func main() {
	TestFile := "hash1.go"
	infile, inerr := os.Open(TestFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("md5: %x  %s\n", md5h.Sum([]byte("")), TestFile)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("sha1: %x  %s\n", sha1h.Sum([]byte("")), TestFile)
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}
}
