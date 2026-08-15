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
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[right] { 
			left = mid + 1
		}else{
			right = mid
		}
	}

	return -1
}
