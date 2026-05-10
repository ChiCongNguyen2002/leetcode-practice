// import "slices"

// func hasDuplicate(nums []int) bool {
//     slices.Sort(nums)

//     count :=0
//     for i:=1 ;i < len(nums);i++ {
//         if nums[i] == nums[i- 1] {
//             count++
//         }
//     }

//     if count > 0 {
//         return true
//     }

//     return false
// }

func hasDuplicate(nums []int) bool {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
        //  1 : 1
        // 2: 1
        // 3 : 2
        if freq[num] > 1 {
            return true
        }
    }

    return false
}
