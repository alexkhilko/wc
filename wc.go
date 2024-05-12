package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"bufio"
	"bytes"
	"strconv"
)

var (
	cFlag bool
	lFlag bool
)

func countBytes(r io.Reader, countLines bool) (int, int, error) {
	buf := make([]byte, 4096)
	bytesCount := 0
	lineSep := []byte{'\n'}
	linesCount := 0
	for {
		n, err := r.Read(buf)
		bytesCount += n
		if countLines {
			linesCount += bytes.Count(buf[:n], lineSep)
		}
		if err == io.EOF {
			return bytesCount, linesCount, nil 
		}
		if err != nil {
			return bytesCount, linesCount, err
		}
	}
}

func main() {
	flag.BoolVar(&cFlag, "c", false, "Count the number of bytes in file")
	flag.BoolVar(&lFlag, "l", false, "Count the number of lines in file")
	flag.Parse()
	arguments := flag.Args()
	if len(arguments) != 1 {
		panic("invalid number of arguments")
	}
	filename := arguments[0]
	f, err := os.Open(filename)
	if err != nil {
		panic("failed to open a file")
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	bytes, lines, err := countBytes(reader, lFlag)
	if err != nil {
		panic("failed to count number of bytes")
	}
	result := ""
	if cFlag {
		result += " " + strconv.Itoa(bytes)
	}
	if lFlag {
		result += " " + strconv.Itoa(lines)
	}
	fmt.Println("", result, filename)
}