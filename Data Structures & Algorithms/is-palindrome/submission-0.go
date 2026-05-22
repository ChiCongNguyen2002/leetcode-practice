func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	return validPalinndrome(s)
	
}

func validPalinndrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s) // wasitacaroracatiswas?

	// remove special characters
	var cleaned []rune

	for _, ch := range s {
		// if character is letter   : a - z
		// if chracter is digital : 0 - 9 keep
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleaned = append(cleaned, ch)
		}
	}

	left :=0
	right := len(cleaned) - 1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}
