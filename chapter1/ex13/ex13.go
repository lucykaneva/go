package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("Time without Join:%f", secs)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs = time.Since(start).Seconds()

	fmt.Printf("Time with Join:%f", secs)

}
