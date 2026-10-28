func twoSum(nums []int, target int) []int {
   mapTwoSum := make(map[int]int)

   for i:=0;i<len(nums);i ++ {
	 	result := target - nums[i]
		// result = 4
		if index,found := mapTwoSum[result]; found {
			return []int{index,i}
		}

		mapTwoSum[nums[i]] = i		
		// mapTwoSum[3] = 0	
   }

   return []int{-1,-1}
}

