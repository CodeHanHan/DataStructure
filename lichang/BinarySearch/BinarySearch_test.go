package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	var a []int = []int{1, 3, 5, 6, 8}

	require.Equal(t, -1, BinarySearch([]int{}, 1))
	require.Equal(t, 4, BinarySearch(a, 8))
	require.Equal(t, -1, BinarySearch(a, 4))
}

func TestBinarySearchRecursive(t *testing.T) {
	var a []int = []int{1, 3, 5, 6, 8}

	require.Equal(t, -1, BinarySearchRecursive([]int{}, 1))
	require.Equal(t, 4, BinarySearchRecursive(a, 8))
	require.Equal(t, -1, BinarySearchRecursive(a, 4))
}

func TestBinarySearchFirst(t *testing.T) {
	var a []int = []int{1, 2, 2, 3, 6, 8, 8}

	require.Equal(t, -1, BinarySearchFirst([]int{}, 1))
	require.Equal(t, 1, BinarySearchFirst(a, 2))
	require.Equal(t, 5, BinarySearchFirst(a, 8))
	require.Equal(t, 1, BinarySearchFirst([]int{1, 2, 2, 2, 3}, 2))
	require.Equal(t, -1, BinarySearchFirst([]int{1, 2, 2, 2, 3}, 4))
}

func TestBinarySearchLast(t *testing.T) {
	var a []int = []int{1, 2, 2, 3, 6, 8, 8}

	require.Equal(t, -1, BinarySearchLast([]int{}, 1))
	require.Equal(t, 2, BinarySearchLast(a, 2))
	require.Equal(t, 6, BinarySearchLast(a, 8))
	require.Equal(t, -1, BinarySearchLast(a, 10))
}

func TestBinarySearchFirstGT(t *testing.T) {
	var a []int = []int{1, 2, 2, 3, 6, 8, 8}

	require.Equal(t, -1, BinarySearchFirstGT([]int{}, 1))
	require.Equal(t, 4, BinarySearchFirstGT(a, 4))
	require.Equal(t, 5, BinarySearchFirstGT(a, 8))
	require.Equal(t, -1, BinarySearchFirstGT(a, 10))
}
