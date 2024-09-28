package matrix

import (
	"reflect"
	"testing"
)

// TestSize calls Matrix's Size() module with a 3x3 Matrix. !
func TestSize(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	matricies := []Matrix{
		{[]float64{1, 2, 3, 4}, []int{4}},
		{[]float64{1, 2, 3, 4}, []int{4, 1}},
		{[]float64{1, 2, 3, 4}, []int{2, 2}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8}, []int{2, 2, 2}},
	}

	coordinateGroups := [][]int{
		{3},
		{2, 0},
		{1, 0},
		{0, 1, 0},
	}

	expecteds := []float64{4, 3, 3, 3}

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
	t.Parallel()
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
	t.Parallel()
	A := Matrix{
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{3, 3},
	}

	B := Matrix{
		[]float64{9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{3, 3},
	}

	C := A.Add(&B)
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

// TestSet calles Matrix.Set with multidimensional coordinates 
func TestSet(t *testing.T) {
	t.Parallel()
	matrix := Matrix{
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{3, 3},
	}

	tests := []struct {
		coords   []int
		value    float64
		expected []float64
	}{
		{[]int{0, 0}, 10, []float64{10, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 1}, 20, []float64{10, 2, 3, 4, 20, 6, 7, 8, 9}},
		{[]int{2, 2}, 30, []float64{10, 2, 3, 4, 20, 6, 7, 8, 30}},
	}

	for _, test := range tests {
		matrix.Set(test.coords, test.value)
		if !reflect.DeepEqual(matrix.data, test.expected) {
			t.Fatalf(`Failed to set value at %v. Expected %v, got %v`,
				test.coords, test.expected, matrix.data)
		}
	}
}

// TestMultiply calles Matrix.Multiply with multidimensional coordinates
func TestMultiply(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name string
		A *Matrix
		B *Matrix
		want *Matrix
	}{
		{
			"2DShouldPass",
			&Matrix{
				[]float64{1, 2, 3, 4, 5, 6},
				[]int{2, 3},
			},
			&Matrix{
				[]float64{7, 10, 13},
				[]int{3, 1},
			},
			&Matrix{
				[]float64{66, 156},
				[]int{2, 1},
			},
		},
		{
			"IdentityMatrix",
			&Matrix{
				[]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
				[]int{3, 3},
			},
			&Matrix{
				[]float64{7, 10, 13, 2, 90, 6},
				[]int{3, 2},
			},
			&Matrix{
				[]float64{7, 10, 13, 2, 90, 6},
				[]int{2, 1},
			},
		},
	}

	for _, test := range tests {
		result := test.A.Multiply(test.B)
		if !reflect.DeepEqual(result.data, test.want.data) {
			t.Fatalf("\nTest case FAILED: %s\nExpected\n%s, got\n%s",
				test.name, test.want.String(), result.String())
		}
	}
}
