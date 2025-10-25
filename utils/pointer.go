package utils

// generic helper
func Ptr[T any](v T) *T {
	return &v
}

func Val[T any](ptr *T) T {
	var zero T
	if ptr == nil {
		return zero
	}
	return *ptr
}
