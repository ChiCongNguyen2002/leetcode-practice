func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{-1,-1}
	}

	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left] + numbers[right] == target {
			return []int{numbers[left],numbers[right]}
		}else if numbers[left] + numbers[right] > target {
			right--
		}else{
			left++
		}
	}

	return []int{-1,-1}
}
