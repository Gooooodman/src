package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("backup_db.sh")
	outputFile, _ := os.OpenFile("backup_db_T.sh", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string([]byte(inputString)[:]) + "\n"
		//fmt.Println(outputString)
		_, err := outputWriter.WriteString(outputString)
		outputWriter.Flush()
		//fmt.Println(n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}
