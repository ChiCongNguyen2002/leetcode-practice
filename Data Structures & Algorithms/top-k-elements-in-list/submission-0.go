func topKFrequent(nums []int, k int) []int {
	// constraints: len(nums) >= 1
	// k >= 1
	// arr: not sorted
	
	// find element frequent >= k
	// input : 1,2,2,3,3,3 k = 2
	// output: [1,2]

	// the mindset presentation: 
	// [1: 1]
	// [2: 2]
	// [3: 3] 
	// k := 2
	// output : find elemenet >= 2 freq 
	// [2 , 3] codition freq >= k

	// step 1: map save freq element
	// step 2: save array count freq 
	// step 3: compare freq map : with k 
	// step 4: output freq >= k 

	type Freq struct {
		val int
		freq int
	}

	// step 1
	mapFreq := make(map[int]int)
	for _,num := range nums {
		mapFreq[num]++ 
	}

	// step 2

	var result []Freq
	for num, freq := range mapFreq {
		result = append(result, Freq{
			val : num,
			freq: freq,
		})
	} 
	// result :
	// [1,1], [2,2] [3,3]
	
	// sort freq 
	sort.Slice(result, func (i, j int) bool{
		return result[i].freq > result[j].freq
	})
	// [3,3] , [2,2], [1,1]

	for i:=0;i<len(result);i++ {
		fmt.Println("result",result[i])
	}

	var output[]int
	for i:=0;i<k;i++ {
		output = append(output, result[i].val)
	}

	return output
}

