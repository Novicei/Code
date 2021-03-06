package _0211018

func intToRoman(num int) string {
	hash := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}
	index := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	res := ""
	for num > 0 {
		for i := len(index) - 1; i >= 0; i-- {
			if num >= index[i] {
				res += hash[index[i]]
				num -= index[i]
				break
			}
		}
	}
	return res
}
