package utils

func FlattenSlices[T any](slices [][]T) []T {
	flattened := make([]T, 0)

	for _, slice := range slices {
		flattened = append(flattened, slice...)
	}

	return flattened
}

func ReverseSlice[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
