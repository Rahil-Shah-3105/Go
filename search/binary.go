package search
import logritelog "log"

// Binary search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// This function uses recursive call to itself.
// If a target is found, the index of the target is returned. Else the function return -1 and ErrNotFound.
func Binary(array []int, target int, lowIndex int, highIndex int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Binary(array=%v,target=%v,lowIndex=%v,highIndex=%v)", array, target, lowIndex, highIndex)
	if highIndex < lowIndex || len(array) == 0 {
		return -1, ErrNotFound
	}
	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		logritelog.Default().Printf("[WARN] <<< Exiting Binary(array=%v,target=%v,lowIndex=%v,highIndex=%v)", array, target, lowIndex, highIndex)
		return mid, nil
	}
}

// BinaryIterative search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// Unlike Binary, this function uses iterative method and not recursive.
// If a target is found, the index of the target is returned. Else the function return -1 and ErrNotFound.
func BinaryIterative(array []int, target int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering BinaryIterative(array=%v,target=%v)", array, target)
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
		logritelog.Default().Printf("[DEBUG] BinaryIterative(array=%v,target=%v): mid → %v", array, target, mid)
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BinaryIterative(array=%v,target=%v)", array, target)
	return -1, ErrNotFound
}

// LowerBound returns index to the first element in the range [0, len(array)-1] that is not less than (i.e. greater or equal to) target.
// return -1 and ErrNotFound if no such element is found.
func LowerBound(array []int, target int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering LowerBound(array=%v,target=%v)", array, target)
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] < target {
		logritelog.Default().Printf("[DEBUG] LowerBound(array=%v,target=%v): mid → %v", array, target, mid)
			startIndex = mid + 1
		} else {
			endIndex = mid - 1
			logritelog.Default().Printf("[DEBUG] LowerBound(array=%v,target=%v): endIndex → %v", array, target, endIndex)
		}
	}

	//when target greater than every element in array, startIndex will out of bounds
	if startIndex >= len(array) {
		return -1, ErrNotFound
	}
	logritelog.Default().Printf("[WARN] <<< Exiting LowerBound(array=%v,target=%v)", array, target)
	return startIndex, nil
}

// UpperBound returns index to the first element in the range [lowIndex, len(array)-1] that is greater than target.
// return -1 and ErrNotFound if no such element is found.
func UpperBound(array []int, target int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering UpperBound(array=%v,target=%v)", array, target)
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
		logritelog.Default().Printf("[DEBUG] UpperBound(array=%v,target=%v): mid → %v", array, target, mid)
			endIndex = mid - 1
		} else {
			startIndex = mid + 1
			logritelog.Default().Printf("[DEBUG] UpperBound(array=%v,target=%v): startIndex → %v", array, target, startIndex)
		}
	}

	//when target greater or equal than every element in array, startIndex will out of bounds
	if startIndex >= len(array) {
		return -1, ErrNotFound
	}
	logritelog.Default().Printf("[WARN] <<< Exiting UpperBound(array=%v,target=%v)", array, target)
	return startIndex, nil
}
