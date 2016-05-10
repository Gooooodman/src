package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//循环读取每一行
//标准输入 os.Stdin 和标准输出 os.Stdout，他们的类型都是 *os.File

func main() {
	inputFile, inputError := os.Open("test.go")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	}
}
