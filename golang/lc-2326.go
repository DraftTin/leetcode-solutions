package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	walls := [4]int{n, m, -1, -1}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = -1
		}
	}
	curDir := 0
	row, col := 0, 0
	for head != nil {
		res[row][col] = head.Val
		head = head.Next
		if curDir == 0 && col+directions[curDir][1] == walls[0] {
			curDir = (curDir + 1) % 4
			walls[3] += 1
			if row+1 == walls[1] {
				break
			}
			row += 1
		} else if curDir == 1 && row+directions[curDir][0] == walls[1] {
			curDir = (curDir + 1) % 4
			walls[0] -= 1
			if col-1 == walls[2] {
				break
			}
			col -= 1
		} else if curDir == 2 && col+directions[curDir][1] == walls[2] {
			curDir = (curDir + 1) % 4
			walls[1] -= 1
			if row-1 == walls[3] {
				break
			}
			row -= 1
		} else if curDir == 3 && row+directions[curDir][0] == walls[3] {
			curDir = (curDir + 1) % 4
			walls[2] += 1
			if col+1 == walls[0] {
				break
			}
			col += 1
		} else {
			row += directions[curDir][0]
			col += directions[curDir][1]
		}
	}
	return res
}
