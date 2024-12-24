package utils

func SliceRemove[T any](slice []T, index int) []T {
	copiedSlice := make([]T, 0, len(slice)-1)
	copiedSlice = append(copiedSlice, slice[:index]...)
	copiedSlice = append(copiedSlice, slice[index+1:]...)
	return copiedSlice
}

func SliceInsert[T any](slice []T, index int, val T) []T {
	slice = append(slice, val)
	copy(slice[index+1:], slice[index:])
	slice[index] = val
	return slice
}

func SliceContains[T comparable](slice []T, val T) bool {
	for _, sliceVal := range slice {
		if sliceVal == val {
			return true
		}
	}
	return false
}
