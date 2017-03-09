package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadInfinetly(path string, inputChannel chan string) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	eofSent := false
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF && !eofSent {
			inputChannel <- "EOF"
			eofSent = true
			fmt.Println("eof sent")
		}
		if line != nil {
			inputChannel <- string(line)
		}
	}
}
