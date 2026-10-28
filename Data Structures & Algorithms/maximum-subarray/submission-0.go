func maxSubArray(nums []int) int {
    if len(nums)== 1{
		return nums[0]
	}

	sum := nums[0]
	max := 0

	for i:=1;i<len(nums);i++{
		if sum  + nums[i] > nums[i] {
			sum += nums[i]
		}else{
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}
	}
	return max
}
