func majorityElement(nums []int) []int {
	num1, num2 := 0, 0
	count1, count2 := 0, 0
	for _, num := range nums {
		if num1 == num {
			count1++
		} else if num2 == num {
			count2++
		} else if count1 == 0 {
			num1 = num
			count1 = 1
		} else if count2 == 0 {
			num2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	for _, num := range nums {
		if num1 == num {
			count1++
		} else if num2 == num {
			count2++
		}
	}
	ans := []int{}
	if count1 > len(nums)/3 {
		ans = append(ans, num1)
	}
	if count2 > len(nums)/3 {
		ans = append(ans, num2)
	}
	return ans
}
