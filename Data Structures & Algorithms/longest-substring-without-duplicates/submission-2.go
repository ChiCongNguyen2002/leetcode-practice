func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0 

	mapStr := make(map[byte]int)
	for i:=0;i<len(s);i++ {
		if index,found := mapStr[s[i]]; found {
			if index >= left {
				// skip ký tự đã từng xét có nằm trong chuỗi hiện tại hay ko?
				left = index + 1
			}
		}

		length := i - left + 1
		mapStr[s[i]] = i

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength 
}
