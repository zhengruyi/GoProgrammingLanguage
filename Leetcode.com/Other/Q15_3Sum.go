package Other

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Sort(sort.IntSlice(nums))
	for i, num := range nums {
		// 相同的数字直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//双指针
		j := i + 1
		k := len(nums) - 1
		for ; j < k; {
			threeSum := num + nums[j] + nums[k]
			if threeSum > 0 {
				k --
			}else if threeSum < 0 {
				j++
			}else {
				res = append(res,[]int{num,nums[j],nums[k]})
				j++
				//这里必须添加这个规避措施，不然【0，0，0，0】
				//的结果会变成 【0，0，0】【0，0，0】
				for ; j < k && nums[j] == nums[j-1]; {
					j++
				}
			}
		}

	}
	return res
}