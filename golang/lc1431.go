package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxVal := 0
	for i := 0; i < len(candies); i++ {
		maxVal = max(maxVal, candies[i])
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxVal {
			res[i] = true
		}
	}
	return res
}
