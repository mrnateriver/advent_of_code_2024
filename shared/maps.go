package shared

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Union[K comparable, V any](a, b map[K]V) map[K]V {
	res := make(map[K]V, max(len(a), len(b)))
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] = v
	}
	return res
}

func Intersect[K comparable, V any](a, b map[K]V) map[K]V {
	res := make(map[K]V, min(len(a), len(b)))
	for k := range a {
		if _, exists := b[k]; exists {
			res[k] = a[k]
		}
	}
	return res
}
