package main

import (
	"bufio"
	"log"
	"os"
)

var delims []byte = []byte{'(', ')', '+', '-', '*', '\'', '"', ' ', '&'}
var digits []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
func ReadFileByLines(source_file string) []string {
	var result []string
	file, err := os.Open(source_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		result = append(result, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return result
}

func main() {
	file := "/home/kaden/Programming/Go/whisper/main.ws"
	lines := ReadFileByLines(file)

	for i := range len(lines) {
		Tokenize(lines[i])
	}
}
