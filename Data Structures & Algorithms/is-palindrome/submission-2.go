func isPalindrome(s string) bool {
    s = strings.ToLower(s) // was it a car or a cat i saw ?
	s = strings.TrimSpace(s) // wasitacaroracatisaw?
	// remove ký tự đặc biệt quên r 

	var cleaned []byte
	for i:=0;i<len(s);i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			cleaned = append(cleaned,c)
		}
	}

	return compare(string(cleaned))
}

func compare(s string) bool {
	left := 0 
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

