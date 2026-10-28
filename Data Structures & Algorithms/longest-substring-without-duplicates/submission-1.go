func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0 

	mapStr := make(map[byte]int)
	for i:=0;i<len(s);i++ {
		if index,found := mapStr[s[i]]; found {
			left = index + 1
		}

		length := i - left + 1
		mapStr[s[i]] = i

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength 
}
