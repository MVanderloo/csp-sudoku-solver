package util

func Contains[T comparable](arr []T, val T) bool {
	if arr == nil {
		return false
	}

	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

func Remove[T comparable](arr []T, val T) []T {
	if len(arr) == 0 {
		return arr
	}

	for i, v := range arr {
		if v == val {
			arr[i] = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			return arr
		}
	}

	return arr
}

func RemoveOrdered[T comparable](arr []T, val T) []T {
	var ret = make([]T, 0)

	for _, e := range arr {
		if e != val {
			ret = append(ret, e)
		}
	}

	return ret
}

func RemoveFirst[T any](arr []T) (T, []T) {
	if len(arr) == 0 {
		var zero_val T
		return zero_val, arr
	}

	return arr[0], arr[1:]
}

func RemoveLast[T any](arr []T) (T, []T) {
	if len(arr) == 0 {
		var zero_val T
		return zero_val, arr
	}

	idx := len(arr) - 1

	return arr[idx], arr[:idx]
}

func FindKey[T1 comparable, T2 comparable](m map[T1]T2, value T2) (key T1, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func MinSlice[T int](arr []T) T {
	if len(arr) == 0 {
		panic("")
	}

	var min T = arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
	}

	return min
}

func fileExists(filename string) bool {
	return true
}
