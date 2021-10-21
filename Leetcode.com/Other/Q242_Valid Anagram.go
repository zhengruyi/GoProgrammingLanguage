package Other

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	stats := make([]int,26)
	for _, c := range s {
		stats[c - rune('a')] ++
	}
	for _, c := range t {
		index := c - rune('a')
		if stats[index] == 0 {
			return false
		}
		stats[index]--
	}
	return true
}
