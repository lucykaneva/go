package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
	countLines(f, counts)
	f.Close()

	for line, n := range counts {
		fmt.Printf("%d\t%s\t\n", n, line)
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		s := input.Text()
		if s[len(s)-1:] == "." || s[len(s)-1:] == ";" || s[len(s)-1:] == "," {
			s = s[:len(s)-1]
		}
		counts[s]++
	}
}
