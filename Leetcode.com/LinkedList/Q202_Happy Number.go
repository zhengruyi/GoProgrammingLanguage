package LinkedList
// 类似于验证链表是否有环
func isHappy(n int) bool {
	x := n
	y := n
	for; x > 1; {

		x = cal(x)
		if x == 1 {
			return true
		}

		y = cal(cal(y))
		if y == 1 {
			return true
		}

		if x == y {
			return false
		}
	}
	return true
}

func cal(x int) int {
	sum := 0
	for; x > 0; {
		sum += (x % 10) * (x % 10)
		x = x /10
	}
	return sum
}