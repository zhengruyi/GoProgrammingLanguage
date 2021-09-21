package String

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	res := ""
	// 算法复杂度为O(n^2)
	for i := 0; i < len(s); i++ {
		// 从当前字符为中心字符开始计算
		s1 := checkPalindrome(s,i,i)
		// 没有中心字符，直接进行计算
		s2 := checkPalindrome(s,i,i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}


func checkPalindrome(s string, i int , j int) string {
	length := len(s)
	for i >= 0 && j < length && s[i] == s[j] {
		i--
		j++
	}
	return string(s[i+1:j])
}