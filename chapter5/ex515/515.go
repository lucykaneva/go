package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	min, err := min()
	if err != nil {
		fmt.Printf("Error in min: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Min: %d\n", min)

	max, err := max(1, 2, 3, 4)
	if err != nil {
		fmt.Printf("Error in max: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Max: %d\n", max)

}
func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("no arguments")
	}
	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min, nil
}
func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("no arguments")
	}
	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max, nil
}
