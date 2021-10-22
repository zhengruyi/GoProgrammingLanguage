package main

import "fmt"

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	flag := 1
	dup := xor

	for; (dup & 1) == 0; {
		flag = 1 << flag
		dup = dup >> 1
	}
	n1 := 0
	for _, num := range nums {
		if (num & flag) == 1 {
			n1 ^=  num
		}
	}
	return []int{n1,xor ^ n1}
}
func main(){
	fmt.Println(2&2)
	singleNumber([]int{1,2,1,3,2,5})
}
