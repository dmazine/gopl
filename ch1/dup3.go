package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func countFileLines(filename string, counts map[string]int) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
}

func countFilesLines(filenames []string, counts map[string]int) {
	for _, filename := range filenames {
		countFileLines(filename, counts)
	}
}

func printLinesCount(counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func main() {
	counts := make(map[string]int)

	countFilesLines(os.Args[1:], counts)

	printLinesCount(counts)
}
