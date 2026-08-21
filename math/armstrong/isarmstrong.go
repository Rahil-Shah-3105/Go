// isarmstrong.go
// description: Checks if the given number is armstrong number or not
// details: An Armstrong number is a n-digit number that is equal to the sum of each of its digits taken to the nth power.
// time complexity: O(log(n))
// space complexity: O(1)
// ref: https://mathlair.allfunandgames.ca/armstrong.php
// author: Kavitha J

package armstrong
import logritelog "log"

import (
	"math"
	"strconv"
)

func IsArmstrong(number int) bool {
	logritelog.Default().Printf(">>> Entering IsArmstrong(number=%v)", number)
	var rightMost int
	var sum int = 0
	var tempNum int = number

	// to get the number of digits in the number
	length := float64(len(strconv.Itoa(number)))

	// get the right most digit and break the loop once all digits are iterated
	for tempNum > 0 {
		rightMost = tempNum % 10
		logritelog.Default().Printf("IsArmstrong(number=%v): rightMost → %v", number, rightMost)
		sum += int(math.Pow(float64(rightMost), length))

		// update the input digit minus the processed rightMost
		tempNum /= 10
	}

	logritelog.Default().Printf("<<< Exiting IsArmstrong(number=%v)", number)
	return number == sum
}
