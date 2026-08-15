func isAnagram(s string, t string) bool {
	mapFreq := make(map[rune]int)

	for _, num := range s {
		mapFreq[num]++
		/// jar
		// jam
		// => false
		// j
		// a
		// r vs m different
	}

	for _,num := range t {
		mapFreq[num]--

		if mapFreq[num] < 0 {
			return false
		}
	}

	return true
}
