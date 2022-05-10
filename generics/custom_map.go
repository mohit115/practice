package main

func createMap[K comparable, V any](input []K, f func(K) V) map[K]V {
	customMap := make(map[K]V)
	for _, e := range input {
		customMap[e] = f(e)
	}
	return customMap
}
