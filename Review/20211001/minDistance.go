package _0211001

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
		dp[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			insert := dp[i-1][j] + 1
			delete := dp[i][j-1] + 1
			changed := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				changed++
			}
			dp[i][j] = min(insert, min(delete, changed))
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
