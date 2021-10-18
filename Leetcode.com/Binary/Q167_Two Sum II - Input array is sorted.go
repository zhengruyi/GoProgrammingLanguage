package Binary
//简单的二分法
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers) - 1
	for ; i < j ; {
		if numbers[i] + numbers[j] < target {
			i++
		}else if numbers[i] + numbers[j] > target {
			j--
		}else{
			return []int{i+1,j+1}
		}
	}
	return []int{}
}
