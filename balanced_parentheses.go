package sprint

func BalancedParentheses(input string) bool {
	stack := make([]rune, 0)
	parenthesesMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, char := range input {
		if isOpeningParenthesis(char) {
			stack = append(stack, char)
		} else if isClosingParenthesis(char) {
			if len(stack) == 0 || parenthesesMap[stack[len(stack)-1]] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isOpeningParenthesis(char rune) bool {
	return char == '(' || char == '[' || char == '{'
}

func isClosingParenthesis(char rune) bool {
	return char == ')' || char == ']' || char == '}'
}
