func twoSum(nums []int, target int) []int {
    mapArr := make(map[int]int)

	for i:=0;i<len(nums);i++ {
		result := target - nums[i]
		if index,found := mapArr[result]; found {
			return []int{index,i}
			// save vi tri tim thay va vi tri hien tai 
		}

		mapArr[nums[i]] = i
		// save ket qua tai vi tri nao? 
	}

	return []int{-1,-1}
}
