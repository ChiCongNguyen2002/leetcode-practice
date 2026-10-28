import "slices"

func groupAnagrams(strs []string) [][]string {
	mapStr := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(str)
		slices.Sort(chars)

		key := string(chars) 
		mapStr[key] = append(mapStr[key],str)
	}

	var result[][]string
	for _, temp := range mapStr {
		result = append(result,temp)
	}
	return result
}	
