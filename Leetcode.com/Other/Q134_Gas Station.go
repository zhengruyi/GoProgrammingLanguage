package Other
// 两点 如果从A出发到不了B,那么AB之间的任意一点都不可能到B
//证明:AB之间的C点,到达C点需要gas的和大于cost,而从A到不了B,那么C更不可能到达B
func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	for i,_ := range gas {
		sum += gas[i] - cost[i]
	}
	if sum < 0 {
		return -1
	}
	start := 0
	tank := 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	return start
}