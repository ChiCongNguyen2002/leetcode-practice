func climbStairs(n int) int {
    if n == 1 || n == 2 {
		return n
	}

	result := make([]int,n)
	result[0] = 1
	result[1] = 2
	
	for i:=2 ;i<n;i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n-1]
}
