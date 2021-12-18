package lists

import (
	"constraints"
)

// All returns true if pred(elem) returns true for all elements in list, otherwise false. The Pred function must return a boolean.
func All[T any](pred func(T) bool, list []T) bool {
	for _, v := range list {
		if !pred(v) {
			return false
		}
	}
	return true
}

// Any returns true if pred(elem) returns true for for at least one element in list, otherwise false. The Pred function must return a boolean.
func Any[T any](pred func(T) bool, list []T) bool {
	for _, v := range list {
		if pred(v) {
			return true
		}
	}
	return false
}

// Concat returns a list in which all the sublists have been appended.
func Concat[T any](lists ...[]T) []T {
	var newList []T
	for _, list := range lists {
		newList = append(newList, list...)
	}
	return newList
}

// Delete returns a copy of list where the first element matching t is deleted, if there is such an element.
func Delete[T comparable](list []T, t T) []T {
	for i, v := range list {
		if v == t {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

// DropLast drops the last element of a List.
func DropLast[T comparable](list []T) ([]T, bool) {
	if len(list) == 0 {
		return []T{}, false
	}
	return list[:len(list)-1], true
}

// DropWhile drop elements from list while pred(elem) returns true and returns the remaining list. The Pred function must return a boolean.
func DropWhile[T any](pred func(T) bool, list []T) []T {
	for i, v := range list {
		if !pred(v) {
			return list[i:]
		}
	}
	return []T{}
}

// Duplicate returns a list containing n copies of term t.
func Duplicate[T any](t T, n int) []T {
	var newList []T
	for i := 0; i < n; i++ {
		newList = append(newList, t)
	}
	return newList
}

// Filter returns a list of all elements in list1 for which pred(elem) returns true. The pred function must return a boolean.
func Filter[T any](pred func(T) bool, list []T) []T {
	newList := make([]T, 0)
	for _, v := range list {
		if pred(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterMap calls fun(elem) on successive elements of list. fun must return a booolean and the new value.
func FilterMap[T any](fun func(T) (bool, T), list []T) []T {
	newList := make([]T, 0)
	for _, v := range list {
		if ok, value := fun(v); ok {
			newList = append(newList, value)
		}
	}
	return newList
}

// FlatMap takes a function from Ts to lists of Us, and a list of As (list) and produces a list of Us by applying the function to every element in list and appending the resulting lists.
func FlatMap[T any, U any](fun func(T) []U, list []T) []U {
	newList := make([]U, 0)
	for _, v := range list {
		newList = append(newList, fun(v)...)
	}
	return newList
}

// Flatten teturns a flattened version of lists.
func Flatten[T any](lists [][]T) []T {
	var newList []T
	for _, list := range lists {
		newList = append(newList, list...)
	}
	return newList
}

// FoldL calls fun(t, acc) on successive elements list, starting with acc. fun must return a new accumulator, which is passed to the next call. The function returns the final value of the accumulator.
func FoldL[T any](fun func(T, T) T, acc T, list []T) T {
	for _, v := range list {
		acc = fun(v, acc)
	}
	return acc
}

// FoldR is like FoldL, but the list is traversed from right to left.
func FoldR[T any](f func(T, T) T, acc T, list []T) T {
	for i := len(list) - 1; i >= 0; i-- {
		acc = f(list[i], acc)
	}
	return acc
}

// ForEach calls fun(elem) for each element in List. This function is used for its side effects and the evaluation order is defined to be the same as the order of the elements in the list.
func ForEach[T any](fun func(T), list []T) {
	for _, v := range list {
		fun(v)
	}
}

// Join inserts sep between each element in list. Has no effect on the empty list and on a singleton list.
func Join[T any](sep T, list []T) []T {
	var newList []T
	for i, v := range list {
		if i > 0 {
			newList = append(newList, sep)
		}
		newList = append(newList, v)
	}
	return newList
}

// Last returns the last element in List.
func Last[T any](list []T) (T, bool) {
	if len(list) == 0 {
		var empty T
		return empty, false
	}
	return list[len(list)-1], true
}

// Map takes a function from Ts to Us, and a list of Ts and produces a list of Us by applying the function to every element in the list. This function is used to obtain the return values.
func Map[T any, U any](fun func(T) U, list []T) []U {
	newList := make([]U, len(list))
	for i, v := range list {
		newList[i] = fun(v)
	}
	return newList
}

// MapFoldL calls fun(elem) on successive elements of list, starting with acc. fun must return the value for the new list and a new accumulator, which is passed to the next call. The function returns the new list and the final value of the accumulator.
func MapFoldL[T any, U any](fun func(T, T) (U, T), acc T, list []T) ([]U, T) {
	var newList []U
	for _, v := range list {
		var u U
		u, acc = fun(v, acc)
		newList = append(newList, u)
	}
	return newList, acc
}

// MapFoldR is like MapFoldL, but the list is traversed from right to left.
func MapFoldR[T any, U any](fun func(T, T) (U, T), acc T, list []T) ([]U, T) {
	var newList []U
	for i := len(list) - 1; i >= 0; i-- {
		var u U
		u, acc = fun(list[i], acc)
		newList = append(newList, u)
	}
	return newList, acc
}

// Max returns the first element of List that compares greater than or equal to all other elements of List.
func Max[T constraints.Ordered](list []T) (T, bool) {
	if len(list) == 0 {
		var empty T
		return empty, false
	}
	max := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max, true
}

// Member returns true if t matches some element of List, otherwise false.
func Member[T comparable](t T, list []T) bool {
	for _, v := range list {
		if v == t {
			return true
		}
	}
	return false
}

// Merge returns the sorted list formed by merging all the sublists. All sublists must be sorted before evaluating this function. When two elements compare equal, the element from the sublist with the lowest position is picked before the other element.
func Merge[T comparable](lists [][]T) []T {
	var newList []T

	return newList
}

// Min returns the first element of List that compares less than or equal to all other elements of List.
func Min[T constraints.Ordered](list []T) (T, bool) {
	if len(list) == 0 {
		var empty T
		return empty, false
	}
	min := list[0]
	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min, true
}

// Nth returns the Nth element of List.
func Nth[T any](n int, list []T) (T, bool) {
	if n >= len(list) {
		var empty T
		return empty, false
	}
	return list[n], true
}

// NthTail returns the Nth tail of List, that is, the sublist of list starting at N+1 and continuing up to the end of the list.
func NthTail[T any](n int, list []T) ([]T, bool) {
	if n >= len(list) {
		var empty []T
		return empty, false
	}
	return list[n+1:], true
}

// Partition partitions list into two lists, where the first list contains all elements for which pred(elem) returns true, and the second list contains all elements for which pred(elem) returns false.
func Partition[T any](pred func(T) bool, list []T) ([]T, []T) {
	var left, right []T
	for _, v := range list {
		if pred(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right
}

// Prefix
