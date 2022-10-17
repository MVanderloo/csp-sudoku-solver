package util

func contains[T comparable](arr []T, val T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

func remove[T comparable](arr []T, val T) []T {
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
