package collect

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	return lo.Map(collection, iteratee)
}

// Unique returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
func Unique[T comparable](collection []T) []T {
	return lo.Uniq(collection)
}

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	return lo.Filter(collection, predicate)
}

// Sum sums the values in a collection. If collection is empty 0 is returned.
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	return lo.Sum(collection)
}

// Max searches the maximum value of a collection.
func Max[T constraints.Ordered](collection []T) T {
	return lo.Max(collection)
}

// Split returns an array of elements split into groups the length of size. If array can't be split evenly,
func Split[T any](collection []T, size int) [][]T {
	return lo.Chunk(collection, size)
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse[T any](collection []T) []T {
	return lo.Reverse(collection)
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
func Shuffle[T any](collection []T) []T {
	return lo.Shuffle(collection)
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	return lo.GroupBy(collection, iteratee)
}

// Count counts the number of elements in the collection.
func Count[T comparable](collection []T) (count int) {
	return len(collection)
}

// CountBy counts the number of elements in the collection for which predicate is true.
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	return lo.CountBy(collection, predicate)
}

// Each iterates over elements of collection and invokes iteratee for each element.
func Each[T any](collection []T, iteratee func(item T, index int)) {
	lo.ForEach(collection, iteratee)
}
