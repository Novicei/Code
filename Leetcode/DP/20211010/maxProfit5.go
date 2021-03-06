package _0211010

func maxProfit5(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}
