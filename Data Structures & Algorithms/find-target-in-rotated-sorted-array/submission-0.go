func search(nums []int, target int) int {
	//index :=0
	for i:=0;i<len(nums);i++ {
		if target == nums[i] {
			return i
		}
	}
	return -1
}

// func search(nums []int, target int) int {
	
// }
