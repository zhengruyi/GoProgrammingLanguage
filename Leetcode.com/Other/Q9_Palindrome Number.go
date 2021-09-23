package Other

func isPalindrome(x int) bool {
	// 负数一定不是对称数
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	num := x
	res := 0
	// 正数就是反转过后看是否和原来一样，注意转换过程中的溢出问题
	for ; x > 0; x = x / 10 {
		tmp := res*10 + x%10
		if (tmp-x%10)/10 != res {
			return false
		}
		res = tmp
	}
	return res == num
}
