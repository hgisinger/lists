package lists

func All[T any](f func(T) bool, list []T) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any[T any](f func(T) bool, list []T) bool {
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func Concat[T any](lists ...[]T) []T {
	var newList []T
	for _, list := range lists {
		newList = append(newList, list...)
	}
	return newList
}

func Delete[T comparable](list []T, t T) []T {
	for i, v := range list {
		if v == t {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func DropLast[T comparable](list []T, t T) []T {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == t {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func DropWhile[T any](list []T, f func(T) bool) []T {
	for i, v := range list {
		if !f(v) {
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
