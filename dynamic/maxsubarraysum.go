// maxsubarraysum.go
// description: Implementation of Kadane's algorithm for Maximum Subarray Sum
// reference: https://en.wikipedia.org/wiki/Maximum_subarray_problem
// time complexity: O(n)
// space complexity: O(1)

package dynamic
import logritelog "log"

import "github.com/TheAlgorithms/Go/math/max"
import stdlog "log"
var log = stdlog.Default()

// MaxSubArraySum returns the sum of the maximum subarray in the input array
func MaxSubArraySum(nums []int) int {
	logritelog.Default().Printf(">>> Entering MaxSubArraySum(nums=%v)", nums)
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max.Int(nums[i], currentSum+nums[i])
		maxSum = max.Int(maxSum, currentSum)
	}
		logritelog.Default().Printf("MaxSubArraySum(nums=%v): currentSum → %v", nums, currentSum)
		logritelog.Default().Printf("MaxSubArraySum(nums=%v): maxSum → %v", nums, maxSum)

	logritelog.Default().Printf("<<< Exiting MaxSubArraySum(nums=%v)", nums)
	return maxSum
}
