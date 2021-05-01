package word

import "testing"
/*
  写几个测试例子，测试文件必须以_test.go，并且所有的测试函数都必须以Test函数开头
 */
func TestPalindrome(t *testing.T) {
	if ! IsPalindrome("detartrated") {
		t.Error("IsPalindrome(detartrated) = false")
	}
	if ! IsPalindrome("kayak") {
		t.Error("IsPalindrome(kayak) = false")
	}
}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error("IsPalindrome(palindrome) = true")
	}
}