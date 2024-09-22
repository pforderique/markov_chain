package matrix

import (
	"testing"
)

// TestSize calls Matrix's Size() module with a 3x3 Matrix.
func TestSize(t *testing.T) {
    A := Matrix{
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{3, 3},
	}
	size := A.Size()

	if size != 9 {
		t.Fatalf(`3x3 matrix A has Size() of %d instead of 9`, size)
	}
}

// TestGet calls Matrix.Get with multidimensional coordinates
func TestGetSuccess(t *testing.T) {
	matricies := []Matrix{
		{[]float64{1, 2, 3, 4}, []int{4}},
		{[]float64{1, 2, 3, 4}, []int{2, 2}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8}, []int{2, 2, 2}},
	}

	coordinateGroups := [][]int{
		{3},
		{1, 0},
		{0, 1, 0},
	}

	expecteds := []float64{4, 3, 3}

	for t_idx, matrix := range matricies {
		coors := coordinateGroups[t_idx]
		expected := expecteds[t_idx]
		result := matrix.Get(coors...)
		if result != expected {
			t.Fatalf(`Test case %d failed. Expected %f, got %f`,
				t_idx, expected, result)
		}
	}
}
