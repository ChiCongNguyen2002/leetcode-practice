import "slices"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	// constraint strs length >= 1

	// clear requirements 
	// group anagrams common
	// [abc, cba, dcb]
	// => [[dcb], [abc,cba]]

	// mindset: 
	// step by step
	// the way 1: 
		// mindset using map[key] : list 
		// sort each strs : using sort chacracters
		// act, cat 
		// a , c, t => sort : act using slices.Sort(arr)
		// c,a,t => a,c,t 
		

	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	var result[][]string
	mapAngrams := sortString(strs)


	for _, group := range mapAngrams {
		result = append(result,group)
	}

	return result
}

func sortString(strs[]string) map[string][]string {
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	mapAnagrams := map[string][]string{}
	for _, str := range strs {
							//  str = "eat"
		arr := []rune(str) // eat : 'e' , 'a' , 't'
		slices.Sort(arr) // ['a', 'e', 't']
		key := string(arr) // "aet"
		mapAnagrams[key] =append(mapAnagrams[key],str)
		// map character sort aet : aet : eat,tea,ate 
	}
	
	return mapAnagrams
}