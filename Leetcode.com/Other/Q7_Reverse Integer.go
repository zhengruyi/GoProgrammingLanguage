package Other
func reverse(x int) int {
	flag := int32(1)
	if x < 0 {
		flag = -1
		x = -1 * x
	}
	res := int32(0)
	var tmp int32
	for ; x > 0; x = x / 10{
		tmp = res * 10 + int32(x % 10)
		// 重点在于如何判断溢出，思路是反向计算，如果计算前后因子不想等则肯定溢出
		if (tmp  - int32(x % 10)) /10 != res {
			return 0
		}
		res = tmp
	}
	return int(res * flag)
}
