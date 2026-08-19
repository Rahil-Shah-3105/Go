// rbc.go
// description: Reflected binary code (RBC)
// details:
// The reflected binary code (RBC), also known just as reflected binary (RB) or Gray code after Frank Gray, is an ordering of the binary numeral system such that two successive values differ in only one bit (binary digit). - [RBC](https://en.wikipedia.org/wiki/Gray_code)
// time complexity: O(n)
// space complexity: O(n)
// author(s) [red_byte](https://github.com/i-redbyte)
// see rbc_test.go

package binary
import logritelog "log"

// SequenceGrayCode The function generates an "Gray code" sequence of length n
func SequenceGrayCode(n uint) []uint {
	logritelog.Default().Printf(">>> Entering SequenceGrayCode(n=%v)", n)
	result := make([]uint, 0)
	var i uint
	for i = 0; i < 1<<n; i++ {
		result = append(result, i^(i>>1))
	}
	logritelog.Default().Printf("<<< Exiting SequenceGrayCode(n=%v)", n)
	return result
}
