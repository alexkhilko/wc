package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"bufio"
	"strconv"
	"unicode"
)

var (
	cFlag bool
	lFlag bool
	wFlag bool
)

func countFile(r io.Reader, countLines, countWords bool) (int, int, int, error) {
	buf := make([]byte, 4096)
	bytesCount := 0
	const lineSep = '\n'
	linesCount := 0
	wordCount := 0
	curWord := false
	for {
		n, err := r.Read(buf)
		bytesCount += n
		for _, b := range(buf[:n]) {
			if countLines && b == lineSep {
				linesCount += 1
			}
			if countWords {
				if unicode.IsSpace(rune(b)) {
					curWord = false
				} else {
					if !curWord {
						wordCount++
					}
					curWord = true
				}
			}
		}
		if err == io.EOF {
			return bytesCount, linesCount, wordCount, nil 
		}
		if err != nil {
			return bytesCount, linesCount, wordCount, err
		}
	}
}

func main() {
	flag.BoolVar(&cFlag, "c", false, "Count the number of bytes in file")
	flag.BoolVar(&lFlag, "l", false, "Count the number of lines in file")
	flag.BoolVar(&wFlag, "w", false, "Count the number of words in file")
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
	bytes, lines, words, err := countFile(reader, lFlag, wFlag)
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
	if wFlag {
		result += " " + strconv.Itoa(words)
	}
	fmt.Println("", result, filename)
}