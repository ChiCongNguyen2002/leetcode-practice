// import "slices"

// func longestConsecutive(nums []int) int {
//     if len(nums) == 0 {
//         return 0
//     }

//     countSeq := 1
//     max :=0
//     slices.Sort(nums) // 0 1 1 2 3 4 5 6
//     for i:=1 ;i< len(nums);i ++ {
//         if nums[i] == nums[i-1] + 1 {
//             countSeq++
//         }

//         if countSeq > max {
//             max = countSeq
//         }
//     }

//     return max => On(log n)
// }


func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num]= true
    }

    longest := 0
    for num := range numSet { 
        // time start duyet. 
        // start duyet tu so be nhat' => lon nhat 
        if _,found := numSet[num - 1]; !found {
            length := 1
            for {
                if _,exist := numSet[num + length]; exist {
                    length++
                }else{
                    break
                }
            }

            if length > longest {
                longest = length
            }
        }
    }
    
    return longest
    // space : O(n)
    // time: O(n)
}

// map : [ 3,1,2,4] => [1,2,3,4] 
// [3: 3]


