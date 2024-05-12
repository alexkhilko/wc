package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"bufio"
)

var (
	cFlag bool
)

func countBytes(r io.Reader) (int, error) {
	buf := make([]byte, 4096)
	count := 0
	for {
		n, err := r.Read(buf)
		count += n
		if err == io.EOF {
			return count, nil 
		}
		if err != nil {
			return count, err
		}
	}
}

func main() {
	flag.BoolVar(&cFlag, "c", false, "Count the number of  bytes in file")
	flag.Parse()
	arguments := flag.Args()
	if len(arguments) != 1 {
		panic("invalid number of arguments")
	}
	filename := arguments[0]
	fmt.Println(arguments)
	f, err := os.Open(filename)
	if err != nil {
		panic("failed to open a file")
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	bytes, err := countBytes(reader)
	if err != nil {
		panic("failed to count number of bytes")
	}
	fmt.Println("%d %s", bytes, filename)
}