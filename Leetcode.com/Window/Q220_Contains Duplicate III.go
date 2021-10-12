package Window

import "math"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 || k == 0 {
		return false
	}

	stats := make(map[int]int,len(nums))
	min := math.MinInt32
	for i,n := range nums {
		if len(stats) > k {
			//除以t+1是为了将0..t的数字全部归类到map的一个索引上
			delete(stats,(nums[i-k-1] - min)/(t+1))
		}
		//减去min是为了将数字全部转化成整数,且保证不会互相覆盖
		idx := (n - min)/(t+1)
		//在一个bucket里面的数字相差纸盒肯定小于等于t
		if _, ok := stats[idx]; ok {
			return true
		}

		if t > 0 {
			//查看相邻bucket是否有符合的值
			if v, ok := stats[idx-1]; ok && n - v <= t {
				return true
			}

			if v, ok := stats[idx+1]; ok && v - n <= t {
				return true
			}
		}
		stats[idx] = n
	}
	return false
}
