package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err

		}
	}
	return nil
}

func main() {

	ReadLine("test2.log", func(out string) {
		r, _ := regexp.Compile("T\\d{2}:\\d{2}:\\d{2}.*")
		o := r.ReplaceAllString(out, "")
		r, _ = regexp.Compile("[[:space:]]")
		fmt.Println(strings.Replace(r.ReplaceAllString(o, "&localtime="), "http://sdklog.tj.65.com/1.html?", "", 1))
	})

}
