package Other

//确定数字范围为32为数字，那么统计每一个bit的累计数字，所以sum表示1bit上面的1出现次数
func singleNumberII(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range nums {
			if ((num >> i) & 1) == 1 {
				sum ++
			}
		}
		if sum != 0 {
			sum = sum % 3
			res |=  sum << i
		}
	}
	//因为int表示64位，所以先强制转化32位，这样可能会变成负数，最后再转化成int
	return int(int32(res))
}