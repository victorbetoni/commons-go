package utils

func KeySlice[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
