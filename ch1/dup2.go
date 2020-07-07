package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}

func countFileLines(fileName string, counts map[string]int) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return
	}

	countLines(file, counts)

	file.Close()
}

func countFilesLines(files []string, counts map[string]int) {
	for _, file := range files {
		countFileLines(file, counts)
	}
}

func showDuplicatedLines(counts map[string]int) {
	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
	}
}

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
		showDuplicatedLines(counts)
		return
	}

	countFilesLines(files, counts)
	showDuplicatedLines(counts)
}
