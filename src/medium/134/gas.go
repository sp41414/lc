package gas

func CanCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0

	// we sum up the cost and gas, if the gas is less than the total
	// cost, there is no possible solution.
	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	if totalGas < totalCost {
		return -1
	}

	// since there is a solution,
	// we look for it by going through and checking
	// the index where we actually have gas after moving
	start, tank := 0, 0
	for i := range gas {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	return start
}
