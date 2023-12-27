package chapter2

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	var stack []byte

	for i := range num {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	stack = stack[:len(stack)-k]

	// Build the final number and remove leading zeros
	result := strings.TrimLeft(string(stack), "0")

	if result == "" {
		return "0"
	}
	return result
}
