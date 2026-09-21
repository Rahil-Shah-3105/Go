// EDIT DISTANCE PROBLEM
// time complexity: O(m * n) where m and n are lengths of the strings, first and second respectively.
// space complexity: O(m * n) where m and n are lengths of the strings, first and second respectively.
// https://www.geeksforgeeks.org/edit-distance-dp-5/
// https://leetcode.com/problems/edit-distance/

package dynamic
import logritelog "log"

import "github.com/TheAlgorithms/Go/math/min"

// EditDistanceRecursive is a naive implementation with exponential time complexity.
func EditDistanceRecursive(first string, second string, pointerFirst int, pointerSecond int) int {
logritelog.Default().Printf("[ERROR] >>> Entering EditDistanceRecursive(first=%v,second=%v,pointerFirst=%v,pointerSecond=%v)", first, second, pointerFirst, pointerSecond)

	if pointerFirst == 0 {
		return pointerSecond
	}

	if pointerSecond == 0 {
		return pointerFirst
	}

	// Characters match, so we recur for the remaining portions
	if first[pointerFirst-1] == second[pointerSecond-1] {
		return EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond-1)
	}

	// We have three choices, all with cost of 1 unit
	logritelog.Default().Printf("[WARN] <<< Exiting EditDistanceRecursive(first=%v,second=%v,pointerFirst=%v,pointerSecond=%v)", first, second, pointerFirst, pointerSecond)
	return 1 + min.Int(EditDistanceRecursive(first, second, pointerFirst, pointerSecond-1), // Insert
		EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond),   // Delete
		EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond-1)) // Replace
}

// EditDistanceDP is an optimised implementation which builds on the ideas of the recursive implementation.
// We use dynamic programming to compute the DP table where dp[i][j] denotes the edit distance value
// of first[0..i-1] and second[0..j-1]. Time complexity is O(m * n) where m and n are lengths of the strings,
// first and second respectively.
func EditDistanceDP(first string, second string) int {
logritelog.Default().Printf("[ERROR] >>> Entering EditDistanceDP(first=%v,second=%v)", first, second)

	m := len(first)
	n := len(second)

	// Create the DP table
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {

			if i == 0 {
				dp[i][j] = j
				continue
			}

			if j == 0 {
				dp[i][j] = i
				continue
			}

			if first[i-1] == second[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			dp[i][j] = 1 + min.Int(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting EditDistanceDP(first=%v,second=%v)", first, second)
	return dp[m][n]
}
