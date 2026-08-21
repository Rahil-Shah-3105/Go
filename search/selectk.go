package search
import logritelog "log"

func SelectK(array []int, k int) (int, error) {
	logritelog.Default().Printf(">>> Entering SelectK(array=%v,k=%v)", array, k)
	if k > len(array) {
	logritelog.Default().Printf("<<< Exiting SelectK(array=%v,k=%v)", array, k)
		return -1, ErrNotFound
	}
	return selectK(array, 0, len(array), len(array)-k), nil
}

// search the element which index is idx
func selectK(array []int, l, r, idx int) int {
	logritelog.Default().Printf(">>> Entering selectK(array=%v,l=%v,r=%v,idx=%v)", array, l, r, idx)
	index := partition(array, l, r)
	if index == idx {
		logritelog.Default().Printf(">>> Entering partition(array=%v,l=%v,r=%v)", array, l, r)
		logritelog.Default().Printf("<<< Exiting partition(array=%v,l=%v,r=%v)", array, l, r)
		return array[index]
	}
	if index < idx {
		return selectK(array, index+1, r, idx)
	}
	logritelog.Default().Printf("<<< Exiting selectK(array=%v,l=%v,r=%v,idx=%v)", array, l, r, idx)
	return selectK(array, l, index, idx)
}

func partition(array []int, l, r int) int {
	elem, j := array[l], l+1
	for i := l + 1; i < r; i++ {
		if array[i] <= elem {
			array[i], array[j] = array[j], array[i]
			j++
		}
	}
	array[l], array[j-1] = array[j-1], array[l]
	return j - 1
}
