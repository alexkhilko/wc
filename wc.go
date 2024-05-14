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
	mFlag bool
)

func getFileCounts(file *os.File, countLines, countWords bool) (int, int, int, int, error) {
	const lineSep = '\n'
	bytesCount := 0
	linesCount := 0
	wordCount := 0
	charCount := 0
	curWord := false

	reader := bufio.NewReader(file)
	for {
		r, size, err := reader.ReadRune()
		if err == io.EOF {
			return bytesCount, linesCount, wordCount, charCount, nil 
		}
		if err != nil {
			return bytesCount, linesCount, wordCount, charCount, err
		}
		charCount++
		bytesCount += size
		if countLines && r == lineSep {
			linesCount += 1
		}
		if countWords {
			if unicode.IsSpace(r) {
				curWord = false
			} else {
				if !curWord {
					wordCount++
				}
				curWord = true
			}
		}
	}
}

func main() {
	flag.BoolVar(&cFlag, "c", false, "Count the number of bytes in file")
	flag.BoolVar(&lFlag, "l", false, "Count the number of lines in file")
	flag.BoolVar(&wFlag, "w", false, "Count the number of words in file")
	flag.BoolVar(&mFlag, "m", false, "Count the number of characters in file")
	flag.Parse()
	if !cFlag && !lFlag && !wFlag && !mFlag {
		cFlag = true
		lFlag = true
		wFlag = true
		mFlag = true
	}
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
	bytes, lines, words, chars, err := getFileCounts(f, lFlag, wFlag)
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
	if mFlag {
		result += " " + strconv.Itoa(chars)
	}
	fmt.Println("", result, filename)
}