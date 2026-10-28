func hasDuplicate(nums []int) bool {
    result := make(map[int]int)

	for _,num := range nums {
		result[num] ++
	}

	for _,num := range result {
		if num > 1 {
			return true 
		}
	}

	return false 
}
