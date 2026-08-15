func isAnagram(s string, t string) bool {
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
	return true
}
