package _0211021

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return left
}
