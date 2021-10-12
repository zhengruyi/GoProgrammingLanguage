package Window
//简易版,直接用map进行检索,唯一注意的是用边界性
func containsNearbyDuplicate(nums []int, k int) bool {
	stats := make(map[int]int,len(nums))
	for i, num := range nums {
		if len(stats) > k {
			delete(stats,nums[i-k-1])
		}
		if _,ok := stats[num];ok {
			return true
		}
		stats[num] = num
	}
	return false
}