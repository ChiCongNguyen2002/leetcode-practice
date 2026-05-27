import "slices"

func threeSum(nums []int) [][]int {
	var result[][]int

	slices.Sort(nums) // -4 -1 -1 0 1 2

	for i:=0;i<len(nums) - 2;i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j:=i + 1;j<len(nums) - 1;j++ {
			if j > i+1 && nums[j] == nums[j - 1] {
				continue
			}
			// Với if j > i+1 && nums[j] == nums[j-1]
			// Ý nghĩa:
			// chỉ skip duplicate SAU phần tử đầu tiên của loop j
			// phần tử đầu tiên vẫn phải được dùng
			
			for k:=j + 1;k<len(nums);k++ {
				if k >j+1 && nums[k] == nums[k - 1]{
					continue
				}
				if nums[i] + nums[j] + nums[k] == 0{
					result = append(result,[]int{nums[i],nums[j],nums[k]})
				}
			}
		}
	}

	return result
}
