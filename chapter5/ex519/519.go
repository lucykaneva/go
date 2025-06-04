//Exercise 5.19: Use panic and recover to write a function that contains no return statement yet returns a non-zero value.

package main

import (
	"fmt"
)

func main() {
	result := returnInt(2)
	fmt.Printf("%d\n", result)
}

func returnInt(x int) (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = p.(int)
		}
	}()
	panic(x)
}
