package testfunc

func ToPointer[T any](d T) *T {
	return &d
}
