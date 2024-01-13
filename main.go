package main

import (
	"fmt"
	"tadd/book/chapter3/exercises"
)

func main() {
	s := ")()(())()()))())))("

	fmt.Println(chapter3.ContigParens(s))
}
