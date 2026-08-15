func isAnagram(s string, t string) bool {
	// string s and string t
	// compare len(s) != len(t) 
	
	// sort s and t compare match 
	// match all => return true
	// miss character => return false 
	s = sortString(s)
	t = sortString(t)

	for i:=0;i<len(s);i++ {
		if s[i] != t[i] {
			return false
		}
	}
	
	return true 
}

func sortString(s string) string{
	chars := []string{} // each characters string 
	for _,str := range s {
		chars = append(chars,string(str))
	}
	sort.Strings(chars)
	return strings.Join(chars, "")
}

