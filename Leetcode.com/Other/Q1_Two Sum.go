package Other
func twoSum(nums []int, target int) []int {
	index := make(map[int]int)
	for i,num := range nums {
		if v,ok := index[target - num]; ok{
			return []int{i,v}
		}
		index[num] = i
	}
	return []int{-1,-1}
}
