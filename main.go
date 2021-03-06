package main

func findDuplicate(nums []int) int {

	if len(nums) <= 2{
		return nums[0]
	}

	slow := nums[0]
	fast := nums[nums[0]]

	for; slow != fast; {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0

	for; slow != fast; {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main(){
	findDuplicate([]int{1,3,4,2,2})
}
