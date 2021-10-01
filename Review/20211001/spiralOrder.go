package _0211001

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left <= right && up <= down {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		for i := up + 1; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		if left != right && up != down {
			for i := right - 1; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			for i := down - 1; i > up; i-- {
				res = append(res, matrix[i][left])
			}
		}
		left++
		right--
		up++
		down--
	}
	return res
}
