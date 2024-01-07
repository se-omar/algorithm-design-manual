package main

import (
	"fmt"
	"tadd/book/chapter3"
)

func main() {
	arr := []int{1, 2, 4, 8, 31, 66}

	fmt.Println(chapter3.BinarySearch(arr, 660, 0, len(arr)))
}
