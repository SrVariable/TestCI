package main

import (
	"os"
	"fmt"
)

func add(nums ...int) int {
	total := 0

	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(add(nums...))
	    fmt.Println(add(nums[3:]...))
	fmt.Println(add(nums[:3]...))
	fmt.Println(os.Args)
}
