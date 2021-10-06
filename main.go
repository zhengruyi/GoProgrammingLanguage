package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	i := 0
	j := 0
	k := 0
	for;k < len(s3);k++ {

		if i < len(s1) && s3[k] == s1[i] {
			i++
			continue
		}

		if j < len(s2) && s3[k] == s2[j] {
			j++
			continue
		}

		return false
	}

	return k == len(s3) && i == len(s1) && j == len(s2)
}

func main() {
isInterleave("aabcc","dbbca","aadbbcbcac")
}
