func isValid(s string) bool {
    if len(s) == 0 {
		return false
	}

	var result []rune
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			result = append(result,char)
		}else{
			if len(result) == 0 {
				return false
			}

			temp := result[len(result)-1] // không nhớ lấy ký tự cuối cùng result 
			if isValidParentheses(temp,char) {
				result = result[:len(result)-1]
			}else{
				return false
			}
		}
	}

	return len(result) == 0
}

func isValidParentheses(a,b rune) bool {
	if a == '[' && b == ']' || a == '{' && b == '}' || a == '(' && b == ')' {
		return true
	}

	return false
}
