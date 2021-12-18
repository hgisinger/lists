package lists

// All returns true if pred(elem) returns true for all elements in list, otherwise false. The Pred function must return a boolean.
func All[T any](pred func(T) bool, list []T) bool {
	for _, v := range list {
		if !Pred(v) {
			return false
		}
	}
	return true
}

// Any returns true if pred(elem) returns true for for at least one element in list, otherwise false. The Pred function must return a boolean.
func Any[T any](f func(T) bool, list []T) bool {
	for _, v := range list {
		if f(v) {
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

// DropLast drops the last element of a List. The list is to be non-empty, otherwise the function panic.
func DropLast[T comparable](list []T) []T {
	return list[:len(list)-1]
}

// DropWhile drop elements from list while pred(elem) returns true and returns the remaining list. The Pred function must return a boolean.
func DropWhile[T any](pred func(T) bool, list []T) []T {
	for i, v := range list {
		if !pred(v) {
			return list[:i]
		}
	}
	return []T{}
}

func Duplicate[T any](t T, n int) []T {
	var newList []T
	for i := 0; i < n; i++ {
		newList = append(newList, t)
	}
	return newList
}

func Filter[T any](f func(T) bool, list []T) []T {
	newList := make([]T, 0)
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterMap[T any](f func(T) (bool, T), list []T) []T {
	newList := make([]T, 0)
	for _, v := range list {
		if ok, value := f(v); ok {
			newList = append(newList, value)
		}
	}
	return newList
}

func FlatMap[T any, U any](f func(T) []U, list []T) []U {
	newList := make([]U, 0)
	for _, v := range list {
		newList = append(newList, f(v)...)
	}
	return newList
}

func Flatten[T any](list [][]T) []T {
	var newList []T
	for _, v := range list {
		newList = append(newList, v...)
	}
	return newList
}

func FoldL[T any](f func(T, T) T, acc T, list []T) T {
	for _, v := range list[1:] {
		acc = f(v, acc)
	}
	return acc
}

func FoldR[T any](f func(T, T) T, acc T, list []T) T {
	for i := len(list) - 1; i >= 0; i-- {
		acc = f(list[i], acc)
	}
	return acc
}

func ForEach[T any](f func(T), list []T) {
	for _, v := range list {
		f(v)
	}
}

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

func Last[T any](list []T) T {
	return list[len(list)-1]
}

func Map[T any, U any](f func(T) U, list []T) []U {
	newList := make([]U, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapFoldL[T any, U any](f func(T, T) (U, T), acc T, list []T) ([]U, T) {
	var newList []U
	for _, v := range list {
		var u U
		u, acc = f(v, acc)
		newList = append(newList, u)
	}
	return newList, acc
}

func MapFoldR[T any, U any](f func(T, T) (U, T), acc T, list []T) ([]U, T) {
	var newList []U
	for i := len(list) - 1; i >= 0; i-- {
		var u U
		u, acc = f(list[i], acc)
		newList = append(newList, u)
	}
	return newList, acc
}

// Max

func Member[T comparable](t T, list []T) bool {
	for _, v := range list {
		if v == t {
			return true
		}
	}
	return false
}

// Merge

// Min

func Nth[T any](n int, list []T) T {
	return list[n]
}

func NthTail[T any](n int, list []T) []T {
	return list[n+1:]
}

func Partition[T any](f func(T) bool, list []T) ([]T, []T) {
	var left, right []T
	for _, v := range list {
		if f(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right
}

// Prefix
