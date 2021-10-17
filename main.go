package main

import "math"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 || k == 0 {
		return false
	}

	stats := make(map[int]int,len(nums))
	min := math.MinInt32
	for i,n := range nums {
		if len(stats) > k {
			delete(stats,(nums[i-k-1] - min)/(t+1))
		}
		idx := (n - min)/(t+1)
		if _, ok := stats[idx]; ok {
			return true
		}

		if t > 0 {
			if v, ok := stats[idx-1]; ok && n - v <= t {
				return true
			}

			if v, ok := stats[idx+1]; ok && v - n <= t {
				return true
			}
		}
	}
	return false
}
func main() {
	path := []int{1,2,3,4,5}
	var res [][]int
	res = append(res, append([]int{6}, path... ))
}
