package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxNum := -1
	for i := 0; i < len(candies); i++ {
		maxNum = max(maxNum, candies[i])
	}
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxNum {
			res[i] = true
		}
	}
	return res
}
