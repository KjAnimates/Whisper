package main

import (
	"KjAnimates/whisper/src"
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot read Environment file. Exiting.")
		return
	}

	file := os.Getenv("TEST_FILE")
	fmt.Printf("file: %v\n", file)

	lines := ReadFileByLines(file)

	for i := range len(lines) {
		src.Tokenize(lines[i])
	}
}
