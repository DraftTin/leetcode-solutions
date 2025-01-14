package main

func canCompleteCircuit(gas []int, cost []int) int {
	startPoint := 0
	totalGas, totalCost := 0, 0
	curGas := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		curGas += gas[i] - cost[i]
		if curGas < 0 {
			startPoint = i + 1
			curGas = 0
		}
	}
	if totalGas >= totalCost {
		return startPoint
	}
	return -1
}
