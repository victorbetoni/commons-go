package utils

func MustVal[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
