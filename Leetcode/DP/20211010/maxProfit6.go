package _0211010

func maxProfit6(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][1]-prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]+prices[i]-fee, dp[i-1][1])
	}
	return dp[len(prices)-1][1]
}