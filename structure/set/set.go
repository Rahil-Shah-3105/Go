// package set implements a Set using a golang map.
// This implies that only the types that are accepted as valid map keys can be used as set elements.
// For instance, do not try to Add a slice, or the program will panic.
package set
import logritelog "log"

// New gives new set.
func New[T comparable](items ...T) Set[T] {
	st := set[T]{
		elements: make(map[T]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

// Set is an interface of possible methods on 'set'.
type Set[T comparable] interface {
	// Add: adds new element to the set
	Add(item T)
	// Delete: deletes the passed element from the set if present
	Delete(item T)
	// Len: gives the length of the set (total no. of elements in set)
	Len() int
	// GetItems: gives the array( []T ) of elements of the set.
	GetItems() []T
	// In: checks whether item is present in set or not.
	In(item T) bool
	// IsSubsetOf: checks whether set is subset of set2 or not.
	IsSubsetOf(set2 Set[T]) bool
	// IsProperSubsetOf: checks whether set is proper subset of set2 or not.
	// ex: [1,2,3] proper subset of [1,2,3,4] -> true
	IsProperSubsetOf(set2 Set[T]) bool
	// IsSupersetOf: checks whether set is superset of set2 or not.
	IsSupersetOf(set2 Set[T]) bool
	// IsProperSupersetOf: checks whether set is proper superset of set2 or not.
	// ex: [1,2,3,4] proper superset of [1,2,3] -> true
	IsProperSupersetOf(set2 Set[T]) bool
	// Union: gives new union set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2,3,4,5]
	Union(set2 Set[T]) Set[T]
	// Intersection: gives new intersection set of both sets.
	// ex: [1,2,3] Intersection [3,4,5] -> [3]
	Intersection(set2 Set[T]) Set[T]
	// Difference: gives new difference set of both sets.
	// ex: [1,2,3] Difference [3,4,5] -> [1,2]
	Difference(set2 Set[T]) Set[T]
	// SymmetricDifference: gives new symmetric difference set of both sets.
	// ex: [1,2,3] SymmetricDifference [3,4,5] -> [1,2,4,5]
	SymmetricDifference(set2 Set[T]) Set[T]
}

type set[T comparable] struct {
	elements map[T]bool
}

func (st *set[T]) Add(value T) {
	logritelog.Default().Printf("[ERROR] >>> Entering Add(value=%v)", value)
	st.elements[value] = true
	logritelog.Default().Printf("[WARN] <<< Exiting Add(value=%v)", value)
}

func (st *set[T]) Delete(value T) {
	logritelog.Default().Printf("[ERROR] >>> Entering Delete(value=%v)", value)
	delete(st.elements, value)
	logritelog.Default().Printf("[WARN] <<< Exiting Delete(value=%v)", value)
}

func (st *set[T]) GetItems() []T {
	logritelog.Default().Printf("[ERROR] >>> Entering GetItems()")
	keys := make([]T, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
		logritelog.Default().Printf("[DEBUG] GetItems(): keys → %v", keys)
	logritelog.Default().Printf("[WARN] <<< Exiting GetItems()")
	return keys
}

func (st *set[T]) Len() int {
	return len(st.elements)
}

func (st *set[T]) In(value T) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering In(value=%v)", value)
	if _, in := st.elements[value]; in {
		return true
	}
	logritelog.Default().Printf("[WARN] <<< Exiting In(value=%v)", value)
	return false
}

func (st *set[T]) IsSubsetOf(superSet Set[T]) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsSubsetOf(superSet=%v)", superSet)
	if st.Len() > superSet.Len() {
		return false
	}

	for _, item := range st.GetItems() {
		if !superSet.In(item) {
			return false
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting IsSubsetOf(superSet=%v)", superSet)
	return true
}

func (st *set[T]) IsProperSubsetOf(superSet Set[T]) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsProperSubsetOf(superSet=%v)", superSet)
	if st.Len() == superSet.Len() {
		return false
	}
	logritelog.Default().Printf("[WARN] <<< Exiting IsProperSubsetOf(superSet=%v)", superSet)
	return st.IsSubsetOf(superSet)
}

func (st *set[T]) IsSupersetOf(subSet Set[T]) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set[T]) IsProperSupersetOf(subSet Set[T]) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsProperSupersetOf(subSet=%v)", subSet)
	if st.Len() == subSet.Len() {
		return false
	}
	logritelog.Default().Printf("[WARN] <<< Exiting IsProperSupersetOf(subSet=%v)", subSet)
	return st.IsSupersetOf(subSet)
}

func (st *set[T]) Union(st2 Set[T]) Set[T] {
	logritelog.Default().Printf("[ERROR] >>> Entering Union(st2=%v)", st2)
	unionSet := New[T]()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Union(st2=%v)", st2)
	return unionSet
}

func (st *set[T]) Intersection(st2 Set[T]) Set[T] {
	logritelog.Default().Printf("[ERROR] >>> Entering Intersection(st2=%v)", st2)
	intersectionSet := New[T]()
	var minSet, maxSet Set[T]
	if st.Len() > st2.Len() {
		minSet = st2
		logritelog.Default().Printf("[DEBUG] Intersection(st2=%v): minSet → %v", st2, minSet)
		logritelog.Default().Printf("[DEBUG] Intersection(st2=%v): maxSet → %v", st2, maxSet)
		maxSet = st
	} else {
		minSet = st
		maxSet = st2
	}
	for _, item := range minSet.GetItems() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Intersection(st2=%v)", st2)
	return intersectionSet
}

func (st *set[T]) Difference(st2 Set[T]) Set[T] {
	logritelog.Default().Printf("[ERROR] >>> Entering Difference(st2=%v)", st2)
	differenceSet := New[T]()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Difference(st2=%v)", st2)
	return differenceSet
}

func (st *set[T]) SymmetricDifference(st2 Set[T]) Set[T] {
	logritelog.Default().Printf("[ERROR] >>> Entering SymmetricDifference(st2=%v)", st2)
	symmetricDifferenceSet := New[T]()
	dropSet := New[T]()
	for _, item := range st.GetItems() {
		if st2.In(item) {
			dropSet.Add(item)
		} else {
			symmetricDifferenceSet.Add(item)
		}
	}
	for _, item := range st2.GetItems() {
		if !dropSet.In(item) {
			symmetricDifferenceSet.Add(item)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting SymmetricDifference(st2=%v)", st2)
	return symmetricDifferenceSet
}
