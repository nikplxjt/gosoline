package funk

func MergeMaps[K comparable, V any, M ~map[K]V](m ...M) (out M) {
	var length int
	for _, item := range m {
		length += len(item)
	}

	out = make(M, length)
	for _, item := range m {
		for k, v := range item {
			out[k] = v
		}
	}

	return
}
