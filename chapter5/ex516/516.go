package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	s, err := stringJoin("/", "abc", "dfghj")
	if err != nil {
		fmt.Printf("Error in string join: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Joined: %s\n", s)
}

func stringJoin(sep string, strings ...string) (string, error) {
	var joined string
	if len(strings) == 0 {
		return "", errors.New("no strings to join")
	}
	for _, s := range strings {
		if joined != "" {
			joined += sep
		}
		joined += s
	}
	return joined, nil
}
