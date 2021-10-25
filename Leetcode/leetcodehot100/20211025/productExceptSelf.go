package _0211025

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	cnt := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= cnt
		cnt *= nums[i]
	}
	return res
}
