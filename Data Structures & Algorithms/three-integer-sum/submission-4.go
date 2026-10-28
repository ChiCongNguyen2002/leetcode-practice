func threeSum(nums []int) [][]int {
	sort.Ints(nums) // [-4, -1, -1, 0, 1, 2]
	var result[][]int

	for i:=0;i<len(nums)- 2 ; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// tim 2 so con lai 

		left := i +1
		right := len(nums) - 1
		for left < right { 
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left ++
			}else if sum > 0 {
				right --
			}else{
				result = append(result,[]int{nums[i], nums[left], nums[right]})
				left ++
				right --

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1]{
					right--
				}
			}
		}
	}

	return result
}
