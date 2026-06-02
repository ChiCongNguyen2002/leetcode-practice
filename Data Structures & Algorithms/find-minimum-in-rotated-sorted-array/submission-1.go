// func findMin(nums []int) int {
// 	slices.Sort(nums)
// 	return nums[0]
// }

func findMin(nums []int) int {
	// 3 4 5 6 1 2
	// 4 5 0 1 2 3

	left :=0 
	right := len(nums) - 1

	for left < right {
		mid := left + ( right - left)/ 2 // mid = 3

		if nums[mid] > nums[right] {
			left = mid + 1 // minimun nằm bên phải mid [mid ... right + 1]
		}else{
			right = mid // minimun nằm bên trái [left ... right ]
		}
	}

	// Một dấu hiệu cho thấy binary search có bug là khi bạn viết: right = mid - 1
	// left = mid + 1
	// "Liệu mid có thể chính là đáp án không?"
	// nếu có thể thì right = mid

	return nums[left]
}
