package utils

func WrapInt(v int) *int {
	var wrap *int
	wrap = new(int)
	wrap = &v
	return wrap
}
