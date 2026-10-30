func isValid(s string) bool {
	openParantheses := "[{("
    stack := []rune{}
	if len(s) % 2 != 0 {
		return false
	}

	for _, char := range s {
		if strings.ContainsRune(openParantheses, char) {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			var expectedChar rune
			if stack[len(stack) - 1] == '{' {
				expectedChar = '}'
			} else if stack[len(stack) - 1] == '[' {
				expectedChar = ']'
			} else {
				expectedChar = ')'
			}

			if char != expectedChar {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	return len(stack) == 0
}