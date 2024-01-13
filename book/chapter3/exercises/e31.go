// A common problem for compilers and text editors is determining whether
// the parentheses in a string are balanced and properly nested. For example, the
// string ((())())() contains properly nested pairs of parentheses, while the strings
// )()( and ()) do not. Give an algorithm that returns true if a string contains
// properly nested and balanced parentheses, and false if otherwise. For full credit,
// identify the position of the first offending parenthesis if the string is not properly
// nested and balanced.

package chapter3

func NestedParens(s string) bool {
	stk := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = append(stk, s[i])
		} else if len(stk) > 0 {
			stk = stk[:len(stk)-1]
		} else {
			return false
		}
	}

	if len(stk) != 0 {
		return false
	}

	return true
}
