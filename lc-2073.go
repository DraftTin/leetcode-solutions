package main

func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for i := 0; i <= k; i++ {
		ans += min(tickets[i], tickets[k])
	}
	for i := k + 1; i < len(tickets); i++ {
		ans += min(tickets[i], tickets[k]-1)
	}
	return ans
}
