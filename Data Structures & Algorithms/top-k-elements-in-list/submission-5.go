func topKFrequent(nums []int, k int) []int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num]++
		// 1 : 1
		// 2 : 2
		// 3 : 3
	}

	var output[]int
	for index, _ := range result {
        output = append(output, index)
    }
	sort.Slice(output, func(i, j int) bool {
        return result[output[i]] > result[output[j]]
    })

	return output[:k]
}
