package utils

type NumLike interface {
	~int | ~float64
}

func AbsDiffNum[T NumLike](x T, y T) T {
	if x < y {
		return y - x
	}
	return x - y
}

func SliceSumNum[T NumLike](slice []T) T {
	var sum T = 0
	for _, val := range slice {
		sum += val
	}
	return sum
}
