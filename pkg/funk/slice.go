package funk

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"math"
)

func Map[S ~[]T1, T1, T2 any, F func(T1) T2](sl S, op F) []T2 {
	result := make([]T2, len(sl))
	for idx, item := range sl {
		result[idx] = op(item)
	}

	return result
}

func Filter[S ~[]T, T any](sl S, pred func(T) bool) []T {
	if len(sl) == 0 {
		return []T{}
	}

	result := make([]T, 0, len(sl))
	for _, item := range sl {
		if pred(item) {
			result = append(result, item)
		}
	}

	return slices.Clip(result)
}

func Reduce[S ~[]T1, T1, T2 any](sl S, op func(T2, T1, int) T2, init T2) T2 {
	result := init

	for idx, item := range sl {
		result = op(result, item, idx)
	}

	return result
}

func ToMap[S ~[]T, T, V any, K comparable](sl S, keyer func(T) (K, V)) map[K]V {
	out := make(map[K]V, len(sl))

	for _, item := range sl {
		k, v := keyer(item)
		out[k] = v
	}

	return out
}

func ToSet[S ~[]T, T comparable](sl S) Set[T] {
	result := make(Set[T], len(sl))

	for _, item := range sl {
		result.Set(item)
	}

	return result
}

func FindFirst[S ~[]T, T comparable](sl S, el T) (out T, ok bool) {
	for _, item := range sl {
		if item == el {
			return item, true
		}
	}

	return
}

func FindFirstFunc[S ~[]T, T any](sl S, pred func(T) bool) (out T, ok bool) {
	for _, item := range sl {
		if pred(item) {
			return item, true
		}
	}

	return
}

func First[S ~[]T, T any](sl S) (out T, ok bool) {
	if len(sl) >= 1 {
		return sl[0], true
	}

	return
}

func Chunk[S ~[]T, T any](sl S, size int) [][]T {
	if size < 1 || len(sl) == 0 {
		return nil
	}

	chunkCount := len(sl) / size
	result := make([][]T, chunkCount)
	for i := 0; i < chunkCount; i++ {
		result[i] = make([]T, 0, size)
	}

	if rem := len(sl) % size; rem != 0 {
		result = append(result, make([]T, 0, rem))
	}

	for i := 0; i < len(sl); i++ {
		result[i/size] = append(result[i/size], sl[i])
	}

	return result
}

func ChunkReduce[S ~[]T, T any](sl S, size int) (out [][]T) {
	length := len(sl)

	if size < 1 || length == 0 {
		return
	}

	return Reduce(sl, func(result [][]T, item T, idx int) [][]T {
		chunkIdx := int(math.Floor(float64(idx) / float64(size)))
		if len(result) < chunkIdx+1 || result[chunkIdx] == nil {
			result = append(result, make([]T, 0, size))
		}

		result[chunkIdx] = append(result[chunkIdx], item)

		return result
	}, make([][]T, 0, length/size))
}

func Reverse[S ~[]T, T any](sl S) (out S) {
	out = make(S, len(sl))
	for i, j := 0, len(sl)-1; i < len(sl) && j >= 0; i, j = i+1, j-1 {
		out[i] = sl[j]
	}

	return
}

func FromSet[T comparable](s Set[T]) (out []T) {
	out = make([]T, 0, len(s))
	for k := range s {
		out = append(out, k)
	}

	return
}

func Uniq[S ~[]T, T comparable](sl S) (out []T) {
	return FromSet(ToSet(sl))
}

func Difference[S ~[]T, T comparable](sl1, sl2 S) (left, right []T) {
	set1, set2 := ToSet(sl1), ToSet(sl2)

	left = make([]T, 0, len(set1))
	right = make([]T, 0, len(set2))

	for _, item := range sl1 {
		if !set2.Contains(item) {
			left = append(left, item)
		}
	}

	for _, item := range sl2 {
		if !set1.Contains(item) {
			right = append(right, item)
		}
	}

	return
}

func Intersect[S ~[]T, T comparable](sl1, sl2 S) []T {
	set2 := ToSet(sl2)
	result := make(Set[T])

	for _, item := range sl1 {
		if set2.Contains(item) {
			result.Set(item)
		}
	}

	return maps.Keys(result)
}
