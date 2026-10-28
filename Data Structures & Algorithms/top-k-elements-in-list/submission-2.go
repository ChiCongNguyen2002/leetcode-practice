func topKFrequent(nums []int, k int) []int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num]++
		// 1 : 1
		// 2 : 2
		// 3 : 3
	}

	type Pair struct {
		num int
		count int
	}

	pairs := make([]Pair,0,len(result))
	for num,count := range result {
		pairs = append(pairs, Pair{
			num : num,
			count: count,
		})
	}

	sort.Slice(pairs, func(i,j int) bool {
		return pairs[i].count > pairs[j].count
	})		

	// Take top k
    output := make([]int, 0, k)
    for i := 0; i < k; i++ {
        output = append(output, pairs[i].num)
    }

	return output
}
