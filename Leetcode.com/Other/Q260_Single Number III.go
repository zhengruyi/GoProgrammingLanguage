package Other

func singleNumberIII(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	//flag的是为了找出从右往左的第一个bit位为1的数字
	// xor - 1,会将最右边的1变成0,然后全部反过来
	flag := xor & (^(xor - 1))

	n1 := 0
	n2 := 0
	for _, num := range nums {
		if (num & flag) == 0  {
			n1 ^=  num
		}else{
			n2 ^= num
		}
	}
	return []int{n1,n2}
}
