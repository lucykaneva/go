package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "Hello $name, today is $day!"

	fmt.Println(expand(s, f))
}

func expand(s string, f func(string) string) string {
	var start, end int
	var newString string
	for i := 0; i < len(s); i++ {
		if s[i] == '$' {
			start = i + 1
			for j := i + 1; j < len(s); j++ {
				if !unicode.IsLetter(rune(s[j])) {
					end = j
					break
				}
			}
			replacement := s[start:end]
			newString += f(replacement)

			i = end
		} else {
			newString += string(s[i])
		}
	}
	return newString
}
func f(s string) string {
	if s == "name" {
		return "Lucy"
	}
	if s == "day" {
		return "Wednesday"
	}
	return ""
}
