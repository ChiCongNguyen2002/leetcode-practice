func lengthOfLongestSubstring(s string) int {
	maxLength := 0 

	mapStr := make(map[byte]int)

	left :=0
	for i:=0;i<len(s);i++{
		if index,found := mapStr[s[i]]; found {
			if index >= left {
				left = index + 1
			}
		}

		mapStr[s[i]] = i

		length := i - left + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
