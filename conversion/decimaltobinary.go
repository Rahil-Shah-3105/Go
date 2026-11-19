/*
Author: Motasim
GitHub: https://github.com/motasimmakki
Date: 14-Oct-2021
*/

// This algorithm will convert any Decimal (+ve integer) number to Binary number.
// https://en.wikipedia.org/wiki/Binary_number
// Function receives a integer as a Decimal number and returns the Binary number.
// Supported integer value range is 0 to 2^(31 -1).
// time complexity: O(log(n))
// space complexity: O(1)

package conversion
import logritelog "log"

// Importing necessary package.
import (
	"errors"
	"strconv"
)

// Reverse() function that will take string,
// and returns the reverse of that string.
func Reverse(str string) string {
	logritelog.Default().Printf("[ERROR] >>> Entering Reverse(str=%v)", str)
	rStr := []rune(str)
	for i, j := 0, len(rStr)-1; i < len(rStr)/2; i, j = i+1, j-1 {
		rStr[i], rStr[j] = rStr[j], rStr[i]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Reverse(str=%v)", str)
	return string(rStr)
}

// DecimalToBinary() function that will take Decimal number as int,
// and return its Binary equivalent as a string.
func DecimalToBinary(num int) (string, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering DecimalToBinary(num=%v)", num)
	if num < 0 {
		return "", errors.New("integer must have +ve value")
	}
	if num == 0 {
		return "0", nil
	}
	var result string = ""
	for num > 0 {
		result += strconv.Itoa(num & 1)
		num >>= 1
	}
	logritelog.Default().Printf("[WARN] <<< Exiting DecimalToBinary(num=%v)", num)
	return Reverse(result), nil
}
