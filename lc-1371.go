package main

// 1. 状态压缩：
// 使用状态压缩的方法来记录每个元音字母出现的次数是否为偶数。
// 可以将每个元音出现的次数使用位掩码表示，每个元音（a、e、i、o、u）使用一个二进制位来记录，0 表示偶数次，1 表示奇数次。
// 在遍历字符串时，跟踪当前的状态并将其与之前的状态进行比较，如果某个状态已经出现过，那么从该状态到当前状态之间的子串满足题目条件。
// 2. 前缀和思想：
// 使用前缀和的思想，通过记录每个状态首次出现的位置，我们可以在遍历字符串时，通过查看当前状态与之前的状态是否相同，来计算最长的子串。
// 这样做可以保证在常数时间内判断从某一位置到当前位置的元音字母是否全部为偶数。

func findTheLongestSubstring(s string) int {
	res := 0
	vowels := map[byte]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}
	firstOccurence := map[int]int{0: -1}
	mask := 0
	for i := 0; i < len(s); i++ {
		if _, ok := vowels[s[i]]; ok {
			mask ^= vowels[s[i]]
		}
		if val, ok := firstOccurence[mask]; ok {
			res = max(res, i-val)
		} else {
			firstOccurence[mask] = i
		}
	}
	return res
}
