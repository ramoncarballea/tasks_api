package collections

func Contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Map[S any, T any](slice []S, f func(S) T) []T {
	var result []T
	for _, s := range slice {
		result = append(result, f(s))
	}
	return result
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, s := range slice {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}
