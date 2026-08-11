func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false 
	}
	// map each character s 
	// map each character t
	// compare map s and t 
	// if not match all character => return false 
	// else => return 

	mapStr := make(map[rune]int)
	for _, str:= range s {
		mapStr[str]++
	}
	for _,str := range t {
		mapStr[str]--
	}
	for _, str:= range mapStr {
		if str > 0 {
			return false
		}
	}

	// Time O(n) 
	// Space O(n) 
	return true
}
