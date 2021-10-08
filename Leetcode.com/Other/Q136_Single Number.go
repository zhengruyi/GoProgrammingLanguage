package Other

// 两个相同数字的抑或之和等于0
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}