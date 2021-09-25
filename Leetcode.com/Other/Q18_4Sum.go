package Other

import "sort"

//主要思想在于四个数字和转化成三个数字之和，即确定一个数字之后，将题目转化为在数组中找到
//三数之和等于给定的target
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Sort(sort.IntSlice(nums))

	for i, n1 := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tt := target - n1
		slice := nums[i+1:]
		for j, n2 := range slice {
			if j > 0 && slice[j] == slice[j-1] {
				continue
			}
			k := j + 1
			l := len(slice) - 1
			for ; k < l; {
				threeSum := n2 + slice[k] + slice[l]
				if threeSum < tt {
					k++
				} else if threeSum > tt {
					l--
				} else {
					res = append(res, []int{n1, n2, slice[k], slice[l]})
					k++
					for ; k < l && slice[k] == slice[k-1]; {
						k++
					}
				}
			}

		}
	}
	return res
}
