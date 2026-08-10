func hasDuplicate(nums []int) bool {
    mapUniqueArr := make(map[int]int)
	for _,num := range nums {
		mapUniqueArr[num]++
	}

	for _, arr := range mapUniqueArr {
		if arr > 1 {
			return true
		}
	}

	return false
}
