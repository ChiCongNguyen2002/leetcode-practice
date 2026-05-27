// func maxArea(heights []int) int {
// 	// the way brute  force
// 	if len(heights) == 0 {
// 		return 0
// 	}

// 	sum := 1
// 	temp := 1
// 	max := 0
// 	for i:=0; i < len(heights) - 1; i ++ {
// 		for j:=i+1;j<len(heights);j++ {
// 			if heights[i] > heights[j] {
// 				temp = heights[j]
// 			}else{
// 				temp = heights[i]
// 			}

// 			sum = (j-i) * temp
// 			if sum > max {
// 				max = sum
// 			}
// 		}
// 	}
// 	return max
// }

func maxArea(heights []int) int {
	// the way brute  force
	if len(heights) == 0 {
		return 0
	}

	temp := 1
	max := 0

	left :=0
	right := len(heights) - 1
	for left < right {
		if heights[left] > heights[right] {
			temp = heights[right]
		}else{
			temp = heights[left]
		}

		length := right - left
		temp *= length
		if temp > max {
			max = temp
		}

		if heights[left] > heights[right] {
			right--
		}else{
			left++
		}
	}

	return max	
}
