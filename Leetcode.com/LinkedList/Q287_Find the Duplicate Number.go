package LinkedList

// nums[i]表示入口的指针来自i,下一个节点在nums[i]处,所以nums转化成链表
//然后位难题转化成了寻找链表成环的入口节点
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
