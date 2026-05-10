// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	mapFreq := make(map[rune]int)

// 	for _, num := range s {
// 		mapFreq[num]++
// 		/// jar
// 		// jam
// 		// => false
// 		// j
// 		// a
// 		// r vs m different
// 	}

// 	for _,num := range t {
// 		mapFreq[num]--

// 		if mapFreq[num] < 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return compareString(sortString(s),sortString(t))
}

func sortString(str string) string {
	chars := []rune(str)
	slices.Sort(chars)
	return string(chars)
}

func compareString(s,t string) bool {
	for i:=0;i<len(s);i++{
		if s[i] != t[i] {
			return false
		}
	}
	return true
}