func twoSum(nums []int, target int) []int {
    twoSum := make(map[int]int)

    for i := 0; i < len(nums); i ++ {
        result := target - nums[i] 
        if index,found := twoSum[result]; found {
            return []int {index,i}
        }
        twoSum[nums[i]] = i
    }
    
    return []int{-1,-1}  
}
