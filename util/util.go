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
