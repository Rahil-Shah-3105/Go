package dynamic
import logritelog "log"

import "github.com/TheAlgorithms/Go/math/min"

// OptimalBST returns the minimum cost of constructing a Binary Search Tree
func OptimalBST(keys []int, freq []int, n int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering OptimalBST(keys=%v,freq=%v,n=%v)", keys, freq, n)
	// Initialize DP table with size n x n
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: single key cost
	for i := 0; i < n; i++ {
		dp[i][i] = freq[i]
	}

	// Build the DP table for sequences of length 2 to n
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			dp[i][j] = int(^uint(0) >> 1) // Initialize with a large value
			sum := sum(freq, i, j)

			// Try every key as root and compute cost
			for k := i; k <= j; k++ {
				// Left cost: dp[i][k-1] is valid only if k > i
				var leftCost int
				if k > i {
					leftCost = dp[i][k-1]
				} else {
					leftCost = 0
				}

				// Right cost: dp[k+1][j] is valid only if k < j
				var rightCost int
				if k < j {
					rightCost = dp[k+1][j]
				} else {
					rightCost = 0
				}

				// Total cost for root k
				cost := sum + leftCost + rightCost

				// Update dp[i][j] with the minimum cost
				dp[i][j] = min.Int(dp[i][j], cost)
			}
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting OptimalBST(keys=%v,freq=%v,n=%v)", keys, freq, n)
	return dp[0][n-1]
}

// Helper function to sum the frequencies
func sum(freq []int, i, j int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering sum(freq=%v,i=%v,j=%v)", freq, i, j)
	total := 0
	for k := i; k <= j; k++ {
		total += freq[k]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting sum(freq=%v,i=%v,j=%v)", freq, i, j)
	return total
}
