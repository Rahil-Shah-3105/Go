package math
import logritelog "log"

// Lerp or Linear interpolation
// This function will return new value in 't' percentage  between 'v0' and 'v1'
func Lerp(v0, v1, t float64) float64 {
	logritelog.Default().Printf("[ERROR] >>> Entering Lerp(v0=%v,v1=%v,t=%v)", v0, v1, t)
	// see: https://en.wikipedia.org/wiki/Linear_interpolation
	return (1-t)*v0 + t*v1
}
