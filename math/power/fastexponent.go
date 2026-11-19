package power
import logritelog "log"

// IterativePower is iterative O(logn) function for pow(x, y)
func IterativePower(n uint, power uint) uint {
	logritelog.Default().Printf("[ERROR] >>> Entering IterativePower(n=%v,power=%v)", n, power)
	var res uint = 1
	for power > 0 {
		if (power & 1) != 0 {
			res = res * n
			logritelog.Default().Printf("[DEBUG] IterativePower(n=%v,power=%v): res → %v", n, power, res)
		}

		power = power >> 1
		n *= n
	}
	logritelog.Default().Printf("[WARN] <<< Exiting IterativePower(n=%v,power=%v)", n, power)
	return res
}

// RecursivePower is recursive O(logn) function for pow(x, y)
func RecursivePower(n uint, power uint) uint {
	logritelog.Default().Printf("[ERROR] >>> Entering RecursivePower(n=%v,power=%v)", n, power)
	if power == 0 {
		return 1
	}
	var temp = RecursivePower(n, power/2)
	if power%2 == 0 {
		return temp * temp
	}
	logritelog.Default().Printf("[WARN] <<< Exiting RecursivePower(n=%v,power=%v)", n, power)
	return n * temp * temp
}

// RecursivePower1 is recursive O(n) function for pow(x, y)
func RecursivePower1(n uint, power uint) uint {
	logritelog.Default().Printf("[ERROR] >>> Entering RecursivePower1(n=%v,power=%v)", n, power)
	if power == 0 {
		return 1
	} else if power%2 == 0 {
		return RecursivePower1(n, power/2) * RecursivePower1(n, power/2)
	} else {
		logritelog.Default().Printf("[WARN] <<< Exiting RecursivePower1(n=%v,power=%v)", n, power)
		return n * RecursivePower1(n, power/2) * RecursivePower1(n, power/2)
	}
}
