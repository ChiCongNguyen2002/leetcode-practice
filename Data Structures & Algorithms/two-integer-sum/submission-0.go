func twoSum(nums []int, target int) []int {
	// constrant : len(nums) >= 2
	// arr : not sorted
	// one valid answer
	// response location nums[i] + nums[j] == target 
	// return location not sorted

	// the way 1: 
	// for loop : 2 for nums[i], nums[j] sum = nums[i] + nums[j]
	// O(n^2)
	// Space : O(1)
	// return twoSumTheWay1(nums,target)

	// the way 2:
	// map : save each record for loop
	// compare output: result = target - nums[i]
	// save map 
	// if found record map = result => found location: result, map
	// Complex : O(n)
	// Space : O(n)
	return twoSumTheWay2(nums,target)
}


func twoSumTheWay1(nums []int, target int) []int {
	for i:=0;i<len(nums) - 1;i ++ {
		for j:=i;j<len(nums);j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{-1,-1}
}


func twoSumTheWay2(nums []int, target int) []int {
	mapTwoSum := make(map[int]int)

	for i:=0;i<len(nums);i ++ {
		result := target - nums[i]
		if index,found := mapTwoSum[result]; found {
			return []int{index,i}
		}
		mapTwoSum[nums[i]] = i		
	}

	return []int{-1,-1}
}