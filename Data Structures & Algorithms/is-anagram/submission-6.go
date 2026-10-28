func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapStr := make(map[rune]int)
	for _, str := range s {
		mapStr[str]++
	}

	for _, str := range t {
		mapStr[str]--
		if mapStr[str] < 0 {
			return false
		}
	}
	return true
}
