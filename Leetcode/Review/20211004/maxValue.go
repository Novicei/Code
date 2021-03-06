package _0211004

func maxValue(grid [][]int) int {
	dp:=make([][]int,len(grid)+1)
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int,len(grid[0])+1)
	}
	for i:=1;i<len(dp);i++{
		for j:=1;j<len(dp[i]);j++{
			dp[i][j]=max(dp[i-1][j],dp[i][j-1])+grid[i-1][j-1]
		}
	}
	return dp[len(grid)][len(grid[0])]
}
