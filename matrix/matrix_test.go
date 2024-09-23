package matrix

import (
	"reflect"
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

// TestString calles Matrix.String with multidimensional coordinates
func TestString(t *testing.T) {
	matricies := []Matrix{
		{[]float64{1, 2, 3, 4}, []int{4}},
		{[]float64{1, 2, 3, 4}, []int{2, 2}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8}, []int{2, 2, 2}},
	}

	expecteds := []string{
		"[1 2 3 4]",
		"[[1 2]\n[3 4]]",
		"[[[1 2]\n[3 4]]\n[[5 6]\n[7 8]]]",
	}

	for t_idx, matrix := range matricies {
		expected := expecteds[t_idx]
		result := matrix.String()
		
		if result != expected {
			t.Fatalf(`Test case %d failed. Expected %s, got %s`,
				t_idx, expected, result)
		}
	}
}

// TestAdd calles Matrix.Add with multidimensional coordinates
func TestAdd(t *testing.T) {
	A := Matrix{
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{3, 3},
	}

	B := Matrix{
		[]float64{9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{3, 3},
	}

	C := A.Add(B)
	expected := Matrix{
		[]float64{10, 10, 10, 10, 10, 10, 10, 10, 10},
		[]int{3, 3},
	}

	if !reflect.DeepEqual(C.data, expected.data){
		t.Fatalf(`Expected %s, got %s`, expected.String(), C.String())
	}

	if !reflect.DeepEqual(C.dims, expected.dims) {
		t.Fatalf(`Expected dimensions %v, got %v`, expected.dims, C.dims)
	}
}
