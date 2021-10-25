package Other

func canConstruct(ransomNote string, magazine string) bool {
	stats := make([]int, 26)
	for _, c := range(magazine) {
		stats[c - rune('a')] ++
	}
	for _, c := range(ransomNote) {
		if stats[c - rune('a')] == 0 {
			return false
		}
		stats[c - rune('a')] --
	}
	return true
}
