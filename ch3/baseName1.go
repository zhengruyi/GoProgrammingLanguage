package main

import "fmt"

func baseName1(s string) string  {
	for i := len(s) - 1; i >= 0; i--{
		if s[i] == '/'{
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}
	return s
}

func main() {
	fmt.Println(baseName1("a/b/c.go"))
}