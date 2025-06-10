package utility

func Of[T any](v T) *T {
	return &v
}
