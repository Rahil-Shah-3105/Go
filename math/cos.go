// author(s) [red_byte](https://github.com/i-redbyte)
// see cos_test.go

package math
import logritelog "log"

import "math"

// Cos  returns the cosine of the radian argument x. [See more](https://en.wikipedia.org/wiki/Sine_and_cosine)
// [Based on the idea of Bhaskara approximation of cos(x)](https://math.stackexchange.com/questions/3886552/bhaskara-approximation-of-cosx)
func Cos(x float64) float64 {
	logritelog.Default().Printf(">>> Entering Cos(x=%v)", x)
	tp := 1.0 / (2.0 * math.Pi)
	x *= tp
	x -= 0.25 + math.Floor(x+0.25)
	x *= 16.0 * (math.Abs(x) - 0.5)
	x += 0.225 * x * (math.Abs(x) - 1.0) //Extra precision
	logritelog.Default().Printf("<<< Exiting Cos(x=%v)", x)
	return x
}
