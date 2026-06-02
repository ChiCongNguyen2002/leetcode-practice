// func search(nums []int, target int) int {
// 	//index :=0
// 	for i:=0;i<len(nums);i++ {
// 		if target == nums[i] {
// 			return i
// 		}
// 	}
// 	return -1
// }

func search(nums []int, target int) int {
	left,right :=0,len(nums) -1 
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// left -> mid đang tăng dần
		if nums[left] <= nums[mid] {
			// target nằm trong đoạn đã sort
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			}else{
				left = mid + 1
			}
		}else{
			// mid -> right đang tăng dần
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			}else{
				right = mid - 1
			}
		}
	}

	return -1
}
