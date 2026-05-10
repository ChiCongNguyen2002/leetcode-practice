import "slices"

func hasDuplicate(nums []int) bool {
    slices.Sort(nums)

    count :=0
    for i:=1 ;i < len(nums);i++ {
        if nums[i] == nums[i- 1] {
            count++
        }
    }

    if count > 0 {
        return true
    }

    return false
}
