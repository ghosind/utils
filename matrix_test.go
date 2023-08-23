package utils

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestMatrix(t *testing.T) {
	a := assert.New(t)

	a.DeepEqual(Matrix[int](0, 0), [][]int{})
	a.DeepEqual(Matrix[int](1, 3), [][]int{{0, 0, 0}})
	a.DeepEqual(Matrix[int](2, 2), [][]int{{0, 0}, {0, 0}})
	a.DeepEqual(Matrix[int](3, 1), [][]int{{0}, {0}, {0}})
	a.DeepEqual(Matrix(2, 2, 1, 2, 3, 4), [][]int{{1, 2}, {3, 4}})
	a.DeepEqual(Matrix(2, 2, 1, 2, 3), [][]int{{1, 2}, {3, 0}})
	a.DeepEqual(Matrix(2, 2, 1, 2, 3, 4, 5), [][]int{{1, 2}, {3, 4}})
	a.DeepEqual(Matrix[int](-1, -2), [][]int{})
}
