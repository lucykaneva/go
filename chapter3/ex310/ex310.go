package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	fmt.Print(comma(s))

}
func comma(s string) string {
	var buf bytes.Buffer
	remainder := len(s) % 3
	if remainder > 0 {
		buf.WriteString(s[:remainder+1])
		buf.WriteString(",")
	}
	for i := remainder; i < len(s); i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < len(s) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
