// https://www.hackerrank.com/challenges/pangrams/problem

package chapter2

import "strings"

func pangrams(s string) string {
	arr := [26]int{}

	s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			arr[s[i]-'a']++
		}
	}

	for _, val := range arr {
		if val == 0 {
			return "not pangram"
		}
	}

	return "pangram"

}
