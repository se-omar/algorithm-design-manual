// Give an algorithm that takes a string S consisting of opening and closing
// parentheses, say )()(())()()))())))(, and finds the length of the longest balanced
// parentheses in S, which is 12 in the example above. (Hint: The solution is not
// necessarily a contiguous run of parenthesis from S.)

package main

func ContigParens(s string) int {
	stk := []byte{}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = append(stk, s[i])
		}

		if len(stk) > 0 && s[i] == ')' {
			stk = stk[:len(stk)-1]
			count += 2
		}
	}
	return count
}
