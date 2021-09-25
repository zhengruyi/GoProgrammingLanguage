package main

import "sort"

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
					for ; k < l; {
						if k > 0 && slice[k] == slice[k-1] {
							k++
						}
					}
				}
			}

		}
	}
	return res
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fourSum(nums, target)
}
