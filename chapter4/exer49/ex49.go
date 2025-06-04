package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	freq := make(map[string]int)
	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
	countLines(f, freq)
	f.Close()
	type pair struct {
		word  string
		count int
	}
	var list []pair
	for w, c := range freq {
		list = append(list, pair{w, c})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].count == list[j].count {
			return list[i].word < list[j].word
		}
		return list[i].count > list[j].count
	})

	for _, p := range list {
		fmt.Printf("%s\t%d\n", p.word, p.count)
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
