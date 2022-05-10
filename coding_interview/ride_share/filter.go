package main

func filter[T interface{}](input []T, f func(T) bool) []T {
	var r []T
	for _, i := range input {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}
