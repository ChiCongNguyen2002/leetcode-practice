func productExceptSelf(nums []int) []int {
	// mindset

	countZero :=0
	temp := 1 
	for i:=0;i<len(nums);i++ {
		if nums[i] == 0{
			countZero++  // > 2
		}else{
		temp *= nums[i]
		}
	}	

	var result[]int
	var flag int

	for i:= 0;i<len(nums);i++ {
		if countZero > 1 {
			result = append(result,0)
		}else if countZero == 1 {
			if nums[i] == 0 {
				nums[i] = temp 
			}else{
				nums[i] = 0
			}
			result = append(result,nums[i])
		}else{
			flag = temp / nums[i]
			result = append(result,flag)
		}
	}	

	return result
}
